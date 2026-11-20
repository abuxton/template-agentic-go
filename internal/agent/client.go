package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	maxResponseBodyBytes = 1 << 20
	maxErrorBodyBytes    = 4 << 10
	maxAttempts          = 3
)

type claudeRequest struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	Messages  []message `json:"messages"`
}

type openAIRequest struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type claudeResponse struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
}

type openAIResponse struct {
	Choices []struct {
		Message message `json:"message"`
	} `json:"choices"`
}

func (c *Client) completeClaude(ctx context.Context, prompt string) (string, error) {
	payload := claudeRequest{
		Model:     c.config.Model,
		MaxTokens: 1024,
		Messages:  []message{{Role: "user", Content: prompt}},
	}
	var response claudeResponse
	err := c.postJSON(ctx, c.config.BaseURL+"/v1/messages", payload, &response, func(request *http.Request) {
		request.Header.Set("x-api-key", c.config.APIKey)
		request.Header.Set("anthropic-version", "2023-06-01")
	})
	if err != nil {
		return "", err
	}
	if len(response.Content) == 0 || response.Content[0].Text == "" {
		return "", fmt.Errorf("claude response contained no text")
	}
	return response.Content[0].Text, nil
}

func (c *Client) completeOpenAI(ctx context.Context, prompt string) (string, error) {
	payload := openAIRequest{
		Model:    c.config.Model,
		Messages: []message{{Role: "user", Content: prompt}},
	}
	var response openAIResponse
	err := c.postJSON(ctx, c.config.BaseURL+"/v1/chat/completions", payload, &response, func(request *http.Request) {
		request.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	})
	if err != nil {
		return "", err
	}
	if len(response.Choices) == 0 || response.Choices[0].Message.Content == "" {
		return "", fmt.Errorf("openai response contained no text")
	}
	return response.Choices[0].Message.Content, nil
}

func (c *Client) postJSON(
	ctx context.Context,
	url string,
	payload any,
	response any,
	configure func(*http.Request),
) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
		if err != nil {
			return fmt.Errorf("create request: %w", err)
		}
		request.Header.Set("Content-Type", "application/json")
		configure(request)

		result, err := c.httpClient.Do(request)
		if err != nil {
			return fmt.Errorf("send request: %w", err)
		}

		responseBody, readErr := readResponseBody(result.Body)
		closeErr := result.Body.Close()
		if readErr != nil {
			return readErr
		}
		if closeErr != nil {
			return fmt.Errorf("close response: %w", closeErr)
		}
		if result.StatusCode >= http.StatusOK && result.StatusCode < http.StatusMultipleChoices {
			if err := json.Unmarshal(responseBody, response); err != nil {
				return fmt.Errorf("decode response: %w", err)
			}
			return nil
		}

		if attempt < maxAttempts && retryableStatus(result.StatusCode) {
			if err := waitBeforeRetry(ctx, result.Header.Get("Retry-After"), attempt); err != nil {
				return err
			}
			continue
		}
		return fmt.Errorf("agent API returned %s: %s", result.Status, truncate(responseBody, maxErrorBodyBytes))
	}

	return nil
}

func readResponseBody(body io.Reader) ([]byte, error) {
	responseBody, err := io.ReadAll(io.LimitReader(body, maxResponseBodyBytes+1))
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}
	if len(responseBody) > maxResponseBodyBytes {
		return nil, fmt.Errorf("agent API response exceeds %d bytes", maxResponseBodyBytes)
	}
	return responseBody, nil
}

func retryableStatus(status int) bool {
	return status == http.StatusTooManyRequests || status >= http.StatusInternalServerError
}

func waitBeforeRetry(ctx context.Context, retryAfter string, attempt int) error {
	delay := time.Duration(1<<uint(attempt-1)) * 200 * time.Millisecond
	if seconds, err := strconv.Atoi(strings.TrimSpace(retryAfter)); err == nil && seconds >= 0 {
		delay = time.Duration(seconds) * time.Second
	}

	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return fmt.Errorf("wait before retry: %w", ctx.Err())
	case <-timer.C:
		return nil
	}
}

func truncate(body []byte, limit int) string {
	if len(body) <= limit {
		return string(body)
	}
	return string(body[:limit]) + "...(truncated)"
}
