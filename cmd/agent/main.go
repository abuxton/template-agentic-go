package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/abuxton/template-agentic-go/internal/agent"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: agent <prompt>")
		os.Exit(2)
	}

	config, err := agent.ConfigFromEnvironment()
	if err != nil {
		fmt.Fprintf(os.Stderr, "configuration error: %v\n", err)
		os.Exit(1)
	}
	client, err := agent.NewClient(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "client error: %v\n", err)
		os.Exit(1)
	}

	response, err := client.Complete(context.Background(), strings.Join(os.Args[1:], " "))
	if err != nil {
		fmt.Fprintf(os.Stderr, "agent request failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(response)
}
