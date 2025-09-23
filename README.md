# gofs-mcp

This is a [Model Context Protocol](https://modelcontextprotocol.io) server that provides LLM tools such as Claude Code with direct access to the [gofs](https://gofs.dev) documentation.

The tool fetches the documentation markdown from the gofs-cli/web repository using the Github API.

## Running

### Requirements to build and run

[Go](https://go.dev/), [Claude Code](https://claude.com/product/claude-code)

### Building the MCP server

To run this MCP locally run the following commands

```bash
  git clone https://github.com/gofs-cli/mcp
  cd mcp
  go build -o gofs-mcp

```

### Using the MCP server

Then navigate to the gofs project and run

```bash
  claude mcp add gofs -- /path/to/mcp/executable
```

to add the mcp to the project or run

```bash
  claude mcp add gofs --scope user -- /path/to/mcp/executable
```

so claude can use the mcp in every project under your current user.

## Cache

This server has a caching feature to speed up fetching the documentation as it is static content that doesn't change often. The cache has a TTL of 1 day.

To clear the cache you must explicitly ask claude code to clear it with a prompt such as `clear the gofs cache` or you can clear it manually in the temp folder at

```bash
cd $TMPDIR/gofs-mcp
```
