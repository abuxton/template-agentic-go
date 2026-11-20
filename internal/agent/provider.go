package agent

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Provider string

const (
	ProviderClaude Provider = "claude"
	ProviderOpenAI Provider = "openai"
)

type Config struct {
	Provider Provider
	Model    string
	APIKey   string
	BaseURL  string
}

func ConfigFromEnvironment() (Config, error) {
	provider := Provider(strings.ToLower(strings.TrimSpace(os.Getenv("AGENT_PROVIDER"))))
	if provider == "" {
		provider = ProviderClaude
	}

	var apiKey, baseURL string
	switch provider {
	case ProviderClaude:
		apiKey = os.Getenv("ANTHROPIC_API_KEY")
		baseURL = os.Getenv("ANTHROPIC_BASE_URL")
	case ProviderOpenAI:
		apiKey = os.Getenv("OPENAI_API_KEY")
		baseURL = os.Getenv("OPENAI_BASE_URL")
	default:
		return Config{}, fmt.Errorf("unsupported agent provider %q", provider)
	}

	if apiKey == "" {
		return Config{}, fmt.Errorf("%s API key is not configured", provider)
	}
	if baseURL == "" {
		if provider == ProviderClaude {
			baseURL = "https://api.anthropic.com"
		} else {
			baseURL = "https://api.openai.com"
		}
	}

	model := strings.TrimSpace(os.Getenv("AGENT_MODEL"))
	if model == "" {
		return Config{}, errors.New("AGENT_MODEL is not configured")
	}

	return Config{
		Provider: provider,
		Model:    model,
		APIKey:   apiKey,
		BaseURL:  strings.TrimRight(baseURL, "/"),
	}, nil
}

type Client struct {
	config     Config
	httpClient *http.Client
}

func NewClient(config Config) (*Client, error) {
	if config.Provider != ProviderClaude && config.Provider != ProviderOpenAI {
		return nil, fmt.Errorf("unsupported agent provider %q", config.Provider)
	}
	if config.Model == "" {
		return nil, errors.New("agent model is required")
	}
	if config.APIKey == "" {
		return nil, errors.New("agent API key is required")
	}
	if config.BaseURL == "" {
		return nil, errors.New("agent base URL is required")
	}

	return &Client{
		config:     config,
		httpClient: http.DefaultClient,
	}, nil
}

func (c *Client) Complete(ctx context.Context, prompt string) (string, error) {
	if strings.TrimSpace(prompt) == "" {
		return "", errors.New("prompt is required")
	}

	switch c.config.Provider {
	case ProviderClaude:
		return c.completeClaude(ctx, prompt)
	case ProviderOpenAI:
		return c.completeOpenAI(ctx, prompt)
	default:
		return "", fmt.Errorf("unsupported agent provider %q", c.config.Provider)
	}
}
