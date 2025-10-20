# Compass Protocol

A simple "Hello World" Model Context Protocol (MCP) server implementation in Go, demonstrating the use of the [official MCP Go SDK](https://github.com/modelcontextprotocol/go-sdk).

## About

This project provides a basic MCP server that showcases how to build MCP-compliant servers using Go. The server implements several simple tools:

- **greet**: Greets a person by name
- **ping**: Simple health check that returns "pong"
- **version**: Returns the server version information

## What is MCP?

The Model Context Protocol (MCP) is an open standard that enables seamless integration between AI applications and data sources. It provides a unified way for AI models to securely access data and tools while maintaining user control and privacy.

## Prerequisites

- Go 1.21 or higher
- Git

## Installation

### Install from Release

Download the latest pre-built binary for your platform from the [Releases](https://github.com/brewinski/compass-protocol/releases) page.

### Build from Source

```bash
# Clone the repository
git clone https://github.com/brewinski/compass-protocol.git
cd compass-protocol

# Install dependencies
go mod download

# Build the server
make build

# Or simply
go build -o compass-protocol
```

## Usage

### Running the Server

The server communicates over stdin/stdout, which is the standard transport for MCP servers:

```bash
./compass-protocol
```

Or using the Makefile:

```bash
make run
```

### Testing the Server

Run the test suite:

```bash
make test
```

### Development

```bash
# Run tests
make test

# Run tests with coverage
make test-coverage

# Build the binary
make build

# Clean build artifacts
make clean

# Run linter (requires golangci-lint)
make lint

# Format code
make fmt

# Show help
make help
```

## Available Tools

The server exposes the following MCP tools:

### greet
Greets a person by name.

**Input:**
```json
{
  "name": "string (required) - The name of the person to greet"
}
```

**Output:**
```json
{
  "greeting": "Hello, {name}! Welcome to Compass Protocol MCP Server."
}
```

### ping
Simple health check.

**Output:**
```json
{
  "message": "pong"
}
```

### version
Returns the server version.

**Output:**
```json
{
  "version": "Compass Protocol MCP Server version: {version}"
}
```

## Project Structure

```
.
├── main.go              # Main server implementation
├── go.mod               # Go module dependencies
├── go.sum               # Go module checksums
├── Makefile             # Build and development tasks
├── .gitignore           # Git ignore patterns
├── .goreleaser.yml      # GoReleaser configuration
├── .github/
│   └── workflows/
│       ├── ci.yml       # CI workflow (test and lint)
│       └── release.yml  # Release workflow (build and publish)
└── README.md            # This file
```

## CI/CD

This project uses GitHub Actions for continuous integration and deployment:

- **CI Workflow**: Runs tests and linting on every push and pull request
- **Release Workflow**: Automatically builds and publishes binaries when a new tag is pushed

### Creating a Release

To create a new release:

```bash
# Tag a new version
git tag -a v1.0.0 -m "Release v1.0.0"

# Push the tag to GitHub
git push origin v1.0.0
```

The release workflow will automatically:
1. Build binaries for multiple platforms (Linux, macOS, Windows)
2. Create a GitHub release
3. Upload the binaries as release assets

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Built with the [official MCP Go SDK](https://github.com/modelcontextprotocol/go-sdk)
- Inspired by the Model Context Protocol specification

## Resources

- [MCP Documentation](https://modelcontextprotocol.io)
- [MCP Go SDK](https://github.com/modelcontextprotocol/go-sdk)
- [MCP Specification](https://spec.modelcontextprotocol.io)