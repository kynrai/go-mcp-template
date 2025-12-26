package content

import (
	"context"

	"github.com/kynrai/go-mcp-template/internal/repo"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type LatestImagesInput struct {
	User       string `json:"name" jsonschema:"the name of the user requesting images"`
	Subscribed bool   `json:"subscribed" jsonschema:"whether the user is a subscriber"`
}

type LatestImagesOutput struct {
	Path string `json:"path" jsonschema:"the path to the latest images"`
}

func latestImages(ctx context.Context, req *mcp.CallToolRequest, input LatestImagesInput) (*mcp.CallToolResult, LatestImagesOutput, error) {
	return nil, LatestImagesOutput{Path: repo.Images(input.User, input.Subscribed)}, nil
}
