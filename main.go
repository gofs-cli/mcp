package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofs-cli/mcp/internal/tools"
	"github.com/gofs-cli/mcp/internal/utils"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	l := log.New(os.Stdout, "[gofs-mcp] ", 2)
	l.Println("Getting routes")
	// fetch the routes when the server starts
	data, err := utils.GetRoutes()

	// server closes if getting the routes fails
	if err != nil {
		fmt.Println(err)
		return
	} else {
		utils.Routes = data
	}

	l.Println("Running server")

	server := mcp.NewServer(
		&mcp.Implementation{Name: "gofs-mcp", Version: "v1.0.0"},
		&mcp.ServerOptions{Instructions: "This is an tool for searching and fetching documentation pages about the gofs tool. gofs (Go Full Stack) is a CLI tool to generate opinionated secure projects using golang + templ + htmx."},
	)

	//mcp.AddTool(server, &mcp.Tool{Name: "fuzzy_search_url", Description: "Find the url for a gofs documentation markdown file using a fuzzy search."}, FuzzySearchRoutes)
	mcp.AddTool(server, &mcp.Tool{Name: "GofsDocumentationUrls", Description: "Fetch the available urls and what content they have from the gofs documentation site."}, tools.FetchUrls)
	mcp.AddTool(server, &mcp.Tool{Name: "GofsDocumentationMarkdown", Description: "Fetch gofs documentations markdown from the given Github API URL. All valid Github API urls are listed in the url column of the GofsDocumentationUrls output."}, tools.FetchMarkdown)
	mcp.AddTool(server, &mcp.Tool{Name: "GofsDocumentationCategories", Description: "Fetch all the possible categories that can be used by GofsDocumentationCategoryMarkdown."}, tools.FetchCategories)
	mcp.AddTool(server, &mcp.Tool{Name: "GofsDocumentationCategoryMarkdown", Description: "Fetch all the documentation in a category, all valid categories are listed by the GofsDocumentationCategories."}, tools.FetchCategoryMarkdown)
	mcp.AddTool(server, &mcp.Tool{Name: "ClearCache", Description: "Clear the documentation cache. ONLY USE THIS FUNCTION WHEN EXPLICITLY ASKED."}, tools.ClearCache)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		l.Fatal(err)
	}
}
