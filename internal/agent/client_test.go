package agent

import (
	"context"
	"net/http"
	"net/http/httptest"
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
