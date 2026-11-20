package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	defer result.Body.Close()

	responseBody, err := io.ReadAll(result.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}
	if result.StatusCode < http.StatusOK || result.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("agent API returned %s: %s", result.Status, string(responseBody))
	}
	if err := json.Unmarshal(responseBody, response); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	return nil
}
