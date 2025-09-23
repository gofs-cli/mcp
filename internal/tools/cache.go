package tools

import (
	"context"
	"os"
	"path/filepath"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// the timestamp of the time the page was cached at is stored at the top of the file separated by a newline
// e.g.
//
// 1758037094
// content...

type ClearCacheInput struct{}

type ClearCacheOutput struct {
	Success bool `json:"success" jsonschema:"A boolean for whether the cache clearing succeeded"`
}

func ClearCache(ctx context.Context, req *mcp.CallToolRequest, input ClearCacheInput) (*mcp.CallToolResult, ClearCacheOutput, error) {
	// just delete the entire gofs-mcp folder in $TMPDIR, it will get recreated the next time AddCache is called
	tempPath := filepath.Join(os.TempDir(), "gofs-mcp")

	err := os.RemoveAll(tempPath)

	if err != nil {
		return nil, ClearCacheOutput{Success: false}, err
	}

	return nil, ClearCacheOutput{Success: true}, nil
}
