package content

import (
	"context"

	"github.com/kynrai/go-mcp-template/internal/repo"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type LatestVideosInput struct {
	User       string `json:"name" jsonschema:"the name of the user requesting videos"`
	Subscribed bool   `json:"subscribed" jsonschema:"whether the user is a subscriber"`
}

type LatestVideosOutput struct {
	Path string `json:"path" jsonschema:"the path to the latest videos"`
}

func latestVideos(ctx context.Context, req *mcp.CallToolRequest, input LatestVideosInput) (*mcp.CallToolResult, LatestVideosOutput, error) {
	return nil, LatestVideosOutput{Path: repo.Videos(input.User, input.Subscribed)}, nil
}
