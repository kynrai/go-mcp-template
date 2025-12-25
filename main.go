package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kynrai/go-mcp-template/tools/greeter"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var (
	stdioTransport      bool = false
	httpStreamTransport bool = false
)

func parseFlags() {
	flag.BoolVar(&stdioTransport, "stdio", false, "Set the MCP transport to stdio")
	flag.BoolVar(&httpStreamTransport, "http", false, "Set the MCP transport to HTTP streaming")

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()

	// check ony one transport is selected using exclusive or
	transports := 0
	if stdioTransport {
		transports++
	}
	if httpStreamTransport {
		transports++
	}
	if transports != 1 {
		exitOneTransport()
	}

}

func exitOneTransport() {
	if !flag.Parsed() {
		flag.Parse()
	}
	fmt.Printf("Exactly one transport must be specified\n\n")
	flag.Usage()
	os.Exit(0)
}

func runMCPServer() {
	parseFlags()

	server := mcp.NewServer(&mcp.Implementation{Name: "example", Version: "v1.0.0"}, nil)
	greeter.RegisterTool(server)

	switch {
	case stdioTransport:
		if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
			log.Fatal(err)
		}
	case httpStreamTransport:
		handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
			return server
		}, nil)
		// Start the HTTP server.
		if err := http.ListenAndServe(":3001", handler); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	default:
		exitOneTransport()
	}

}

func main() {
	runMCPServer()
}
