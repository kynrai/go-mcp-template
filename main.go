package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kynrai/go-mcp-template/tools/content"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var (
	MCPHost           string
	MCPPort           string
	MCPTransportStdio bool
	MCPTransportHTTP  bool
)

func parseConfig() {
	flag.StringVar(&MCPHost, "host", "localhost", "The host for the MCP server")
	flag.StringVar(&MCPPort, "port", "3001", "The port for the MCP server")
	flag.BoolVar(&MCPTransportStdio, "stdio", false, "Set the MCP transport to stdio")
	flag.BoolVar(&MCPTransportHTTP, "http", false, "Set the MCP transport to HTTP streaming")

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()

	transports := 0
	if MCPTransportStdio {
		transports++
	}
	if MCPTransportHTTP {
		transports++
	}

	if transports == 0 {
		log.Fatal("No transport specified, defaulting to -stdio ")
	} else if transports > 1 {
		log.Fatal("You can only specify one transport at a time")
	}
}

func runMCPServer() {
	parseConfig()
	server := mcp.NewServer(&mcp.Implementation{Name: "example", Version: "v1.0.0"}, nil)
	content.RegisterTools(server)

	switch {
	case MCPTransportHTTP:
		handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
			return server
		}, nil)
		// Start the HTTP server.
		if err := http.ListenAndServe(":3001", handler); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	default:
		if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	runMCPServer()
}
