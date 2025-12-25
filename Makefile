inspector:
	@go build -o bin/mcp . && bunx @modelcontextprotocol/inspector@latest bin/mcp -stdio
.PHONY: inspector

build:
	@go build -o bin/mcp .
.PHONY: build
