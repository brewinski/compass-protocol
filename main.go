// Copyright 2025 Compass Protocol Authors. All rights reserved.
// Use of this source code is governed by an MIT-style license.

// Package main implements a simple "Hello World" MCP (Model Context Protocol) server
// that demonstrates basic MCP server functionality using the official Go SDK.
package main

import (
	"context"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Version is the application version, set during build time
var Version = "dev"

func main() {
	// Create a new MCP server instance
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "compass-protocol",
		Version: Version,
	}, nil)

	// Define the input structure for the greet tool
	type GreetInput struct {
		Name string `json:"name" jsonschema:"the name of the person to greet"`
	}

	// Add a simple "greet" tool to the server
	mcp.AddTool(server, &mcp.Tool{
		Name:        "greet",
		Description: "Greets a person by name",
	}, func(ctx context.Context, req *mcp.CallToolRequest, input GreetInput) (*mcp.CallToolResult, any, error) {
		greeting := "Hello, " + input.Name + "! Welcome to Compass Protocol MCP Server."

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: greeting},
			},
		}, nil, nil
	})

	// Add a "ping" tool for health checks
	mcp.AddTool(server, &mcp.Tool{
		Name:        "ping",
		Description: "Simple health check that returns 'pong'",
	}, func(ctx context.Context, req *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, any, error) {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "pong"},
			},
		}, nil, nil
	})

	// Add a "version" tool to get server version
	mcp.AddTool(server, &mcp.Tool{
		Name:        "version",
		Description: "Returns the server version",
	}, func(ctx context.Context, req *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, any, error) {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Compass Protocol MCP Server version: " + Version},
			},
		}, nil, nil
	})

	// Log server startup
	log.Printf("Starting Compass Protocol MCP Server v%s", Version)
	log.Println("Server will communicate over stdin/stdout")

	// Run the server over stdin/stdout transport
	// This is the standard transport for MCP servers
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Printf("Server error: %v", err)
		os.Exit(1)
	}
}
