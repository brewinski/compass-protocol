# Examples

This directory contains example code demonstrating how to use the Compass Protocol MCP Server.

## Client Example

The `client.go` file demonstrates how to create an MCP client that connects to the Compass Protocol server and calls its tools.

### Running the Example

1. First, make sure the Compass Protocol server is built:
   ```bash
   cd ..
   make build
   ```

2. Run the example client:
   ```bash
   go run client.go
   ```

   Or if you've installed the server binary to your PATH:
   ```bash
   cd examples
   go run client.go
   ```

### Expected Output

```
Connecting to Compass Protocol MCP Server...
Connected successfully!

Calling 'version' tool...
  Response: Compass Protocol MCP Server version: dev

Calling 'ping' tool...
  Response: pong

Calling 'greet' tool with name 'Alice'...
  Response: Hello, Alice! Welcome to Compass Protocol MCP Server.

All tool calls completed successfully!
```

## Other Integration Examples

### Using with Claude Desktop

To use this MCP server with Claude Desktop, add it to your Claude configuration file:

**macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`

**Windows**: `%APPDATA%/Claude/claude_desktop_config.json`

```json
{
  "mcpServers": {
    "compass-protocol": {
      "command": "/path/to/compass-protocol"
    }
  }
}
```

### Using with Other MCP Clients

The server implements the standard MCP protocol and can be used with any MCP-compatible client. The server communicates over stdin/stdout using JSON-RPC 2.0.

Example JSON-RPC request to list tools:

```bash
echo '{"jsonrpc": "2.0", "id": 1, "method": "tools/list", "params": {}}' | ./compass-protocol
```
