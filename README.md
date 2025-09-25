# gofs-mcp

This is a [Model Context Protocol](https://modelcontextprotocol.io) server that provides LLM tools such as Claude Code with direct access to the [gofs](https://gofs.dev) documentation.

The tool fetches the documentation markdown from the gofs-cli/web repository using the Github API.

## Running

### Requirements to build and run

[Go](https://go.dev/), [Claude Code](https://claude.com/product/claude-code)

### Building the MCP server

To run this MCP locally run the following commands

```bash
  go install github.com/gofs-cli/mcp@latest
```

### Using the MCP server

Then either navigate to the directory of the gofs project you want to use claude code with and run

```bash
  claude mcp add gofs -- $(go env GOPATH)/bin/mcp
```

to add the mcp to only that project or from any directory run

```bash
  claude mcp add gofs --scope user -- $(go env GOPATH)/bin/mcp
```

to allow claude to use the mcp in every directory.

## Cache

This server has a caching feature to speed up fetching the documentation as it is static content that doesn't change often. The cache has a TTL of 1 day.

To clear the cache you must explicitly ask claude code to clear it with a prompt such as `clear the gofs cache` or you can clear it manually in the temp folder at

```bash
cd $TMPDIR/gofs-mcp
```
