package content

import "github.com/modelcontextprotocol/go-sdk/mcp"

func RegisterTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{Name: "latestImages", Description: "get the latest images if the user is subscribed"}, latestImages)
	mcp.AddTool(server, &mcp.Tool{Name: "latestVideos", Description: "get the latest videos if the user is subscribed"}, latestVideos)
}
