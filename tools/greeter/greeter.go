package greeter

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Input struct {
	Name string `json:"name" jsonschema:"the name of the person to greet"`
}

type Output struct {
	Greeting string `json:"greeting" jsonschema:"the greeting to tell to the user"`
}

func sayHi(ctx context.Context, req *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, Output, error) {
	return nil, Output{Greeting: "Hi " + input.Name}, nil
}

func RegisterTool(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{Name: "greet", Description: "say hi"}, sayHi)
}
