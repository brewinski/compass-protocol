// Copyright 2025 Compass Protocol Authors. All rights reserved.
// Use of this source code is governed by an MIT-style license.

// Package main provides a simple example of an MCP client that connects to the
// Compass Protocol server and calls its tools.
package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	ctx := context.Background()

	// Create a new MCP client
	client := mcp.NewClient(&mcp.Implementation{
		Name:    "compass-client",
		Version: "1.0.0",
	}, nil)

	// Connect to the Compass Protocol server
	// The server binary should be in the parent directory (for development)
	// or in your PATH (for production use)
	serverPath := "../compass-protocol" // Development path
	transport := &mcp.CommandTransport{
		Command: exec.Command(serverPath),
	}

	fmt.Println("Connecting to Compass Protocol MCP Server...")
	session, err := client.Connect(ctx, transport, nil)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer session.Close()

	fmt.Println("Connected successfully!")
	fmt.Println()

	// Call the "version" tool
	fmt.Println("Calling 'version' tool...")
	versionResult, err := session.CallTool(ctx, &mcp.CallToolParams{
		Name:      "version",
		Arguments: map[string]any{},
	})
	if err != nil {
		log.Fatalf("Failed to call version tool: %v", err)
	}
	if versionResult.IsError {
		log.Fatal("version tool returned an error")
	}
	for _, content := range versionResult.Content {
		if textContent, ok := content.(*mcp.TextContent); ok {
			fmt.Printf("  Response: %s\n", textContent.Text)
		}
	}
	fmt.Println()

	// Call the "ping" tool
	fmt.Println("Calling 'ping' tool...")
	pingResult, err := session.CallTool(ctx, &mcp.CallToolParams{
		Name:      "ping",
		Arguments: map[string]any{},
	})
	if err != nil {
		log.Fatalf("Failed to call ping tool: %v", err)
	}
	if pingResult.IsError {
		log.Fatal("ping tool returned an error")
	}
	for _, content := range pingResult.Content {
		if textContent, ok := content.(*mcp.TextContent); ok {
			fmt.Printf("  Response: %s\n", textContent.Text)
		}
	}
	fmt.Println()

	// Call the "greet" tool with a name
	fmt.Println("Calling 'greet' tool with name 'Alice'...")
	greetResult, err := session.CallTool(ctx, &mcp.CallToolParams{
		Name: "greet",
		Arguments: map[string]any{
			"name": "Alice",
		},
	})
	if err != nil {
		log.Fatalf("Failed to call greet tool: %v", err)
	}
	if greetResult.IsError {
		log.Fatal("greet tool returned an error")
	}
	for _, content := range greetResult.Content {
		if textContent, ok := content.(*mcp.TextContent); ok {
			fmt.Printf("  Response: %s\n", textContent.Text)
		}
	}
	fmt.Println()

	fmt.Println("All tool calls completed successfully!")
}
