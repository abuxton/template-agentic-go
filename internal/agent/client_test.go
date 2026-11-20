package agent

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCompleteClaude(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/v1/messages" {
			t.Fatalf("path = %q, want /v1/messages", request.URL.Path)
		}
		if request.Header.Get("x-api-key") != "test-key" {
			t.Fatalf("missing Claude API key header")
		}
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte(`{"content":[{"text":"hello from Claude"}]}`))
	}))
	defer server.Close()

	client, err := NewClient(Config{
		Provider: ProviderClaude,
		Model:    "test-model",
		APIKey:   "test-key",
		BaseURL:  server.URL,
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := client.Complete(context.Background(), "hello")
	if err != nil {
		t.Fatal(err)
	}
	if got != "hello from Claude" {
		t.Fatalf("response = %q, want %q", got, "hello from Claude")
	}
}

func TestCompleteOpenAI(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/v1/chat/completions" {
			t.Fatalf("path = %q, want /v1/chat/completions", request.URL.Path)
		}
		if request.Header.Get("Authorization") != "Bearer test-key" {
			t.Fatalf("missing OpenAI authorization header")
		}
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write([]byte(`{"choices":[{"message":{"role":"assistant","content":"hello from OpenAI"}}]}`))
	}))
	defer server.Close()

	client, err := NewClient(Config{
		Provider: ProviderOpenAI,
		Model:    "test-model",
		APIKey:   "test-key",
		BaseURL:  server.URL,
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := client.Complete(context.Background(), "hello")
	if err != nil {
		t.Fatal(err)
	}
	if got != "hello from OpenAI" {
		t.Fatalf("response = %q, want %q", got, "hello from OpenAI")
	}
}

func TestCompleteRejectsEmptyPrompt(t *testing.T) {
	client, err := NewClient(Config{
		Provider: ProviderOpenAI,
		Model:    "test-model",
		APIKey:   "test-key",
		BaseURL:  "https://api.openai.com",
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := client.Complete(context.Background(), " \t"); err == nil {
		t.Fatal("expected empty prompt error")
	}
}

func TestCompleteReturnsProviderError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Retry-After", "0")
		http.Error(writer, "provider unavailable", http.StatusServiceUnavailable)
	}))
	defer server.Close()

	client, err := NewClient(Config{
		Provider: ProviderOpenAI,
		Model:    "test-model",
		APIKey:   "test-key",
		BaseURL:  server.URL,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Complete(context.Background(), "hello")
	if err == nil {
		t.Fatal("expected provider error")
	}
	if got := err.Error(); !containsAll(got, "503", "provider unavailable") {
		t.Fatalf("error = %q, want status and provider message", got)
	}
}

func TestCompleteRetriesTransientProviderFailure(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		attempts++
		if attempts == 1 {
			writer.Header().Set("Retry-After", "0")
			http.Error(writer, "try again", http.StatusBadGateway)
			return
		}
		_, _ = writer.Write([]byte(`{"choices":[{"message":{"role":"assistant","content":"recovered"}}]}`))
	}))
	defer server.Close()

	client, err := NewClient(Config{
		Provider: ProviderOpenAI,
		Model:    "test-model",
		APIKey:   "test-key",
		BaseURL:  server.URL,
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := client.Complete(context.Background(), "hello")
	if err != nil {
		t.Fatal(err)
	}
	if got != "recovered" {
		t.Fatalf("response = %q, want %q", got, "recovered")
	}
	if attempts != 2 {
		t.Fatalf("attempts = %d, want 2", attempts)
	}
}

func TestCompleteRejectsMalformedResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(`{"choices":`))
	}))
	defer server.Close()

	client, err := NewClient(Config{
		Provider: ProviderOpenAI,
		Model:    "test-model",
		APIKey:   "test-key",
		BaseURL:  server.URL,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Complete(context.Background(), "hello")
	if err == nil {
		t.Fatal("expected malformed response error")
	}
}

func containsAll(value string, parts ...string) bool {
	for _, part := range parts {
		if !strings.Contains(value, part) {
			return false
		}
	}
	return true
}
