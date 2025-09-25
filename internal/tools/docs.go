package tools

import (
	"context"
	"slices"
	"strings"

	"github.com/gofs-cli/mcp/internal/utils"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type FetchMarkdownInput struct {
	URL string `json:"url" jsonschema:"the url of the gofs markdown documentation"`
}

type FetchMarkdownOutput struct {
	MarkdownContent string `json:"markdown_content" jsonschema:"markdown from the requested url"`
}

type FetchCategoriesInput struct{}

type FetchCategoriesOutput struct {
	Categories string `json:"categories" jsonschema:"the possible categories than can be used in fetch_documentation_category_markdown"`
}

type FetchCategoryMarkdownInput struct {
	Category string `json:"url" jsonschema:"the name of the category from the fetch_documentation_categories tool"`
}

type FetchCategoryMarkdownOutput struct {
	MarkdownContent string `json:"markdown_content" jsonschema:"markdown from the requested url"`
}

type FetchDocumentationUrlsInput struct{}

type FetchDocumentationUrlsOutput struct {
	UrlsList string `json:"urls_list" jsonschema:"all the documentation paths and their urls"`
}

// type FuzzySearchInput struct {
// 	SearchTerm string `json:"search_term" jsonschema:"the type of document page you want to find the url for"`
// }

// type FuzzySearchOutput struct {
// 	URL string `json:"url" jsonschema:"the url of the document page found by the fuzzy search, this will be empty if no url is found"`
// }

// func FuzzySearchRoutes(ctx context.Context, req *mcp.CallToolRequest, input FuzzySearchInput) (*mcp.CallToolResult, FuzzySearchOutput, error) {
// 	fmt.Println("Fuzzy searching for \"" + input.SearchTerm + "\"")

// 	return nil, FuzzySearchOutput{URL: ""}, nil
// }

func FetchMarkdown(ctx context.Context, req *mcp.CallToolRequest, input FetchMarkdownInput) (*mcp.CallToolResult, FetchMarkdownOutput, error) {
	content, err := utils.FetchSingleMarkdown(input.URL)
	return nil, FetchMarkdownOutput{MarkdownContent: content}, err
}

func FetchUrls(ctx context.Context, req *mcp.CallToolRequest, input FetchDocumentationUrlsInput) (*mcp.CallToolResult, FetchDocumentationUrlsOutput, error) {
	//the routes get loaded when the server starts, so no logic is needed here
	return nil, FetchDocumentationUrlsOutput{UrlsList: utils.FormatRoutes(utils.Routes)}, nil
}

func FetchCategoryMarkdown(ctx context.Context, req *mcp.CallToolRequest, input FetchCategoryMarkdownInput) (*mcp.CallToolResult, FetchCategoryMarkdownOutput, error) {
	var routes []utils.RouteData

	for _, value := range utils.Routes {
		if strings.HasPrefix(value.Path, input.Category) {
			routes = append(routes, value)
		}
	}

	content := ""

	for _, route := range routes {
		markdown, err1 := utils.FetchSingleMarkdown(route.URL)

		if err1 != nil {
			return nil, FetchCategoryMarkdownOutput{}, err1
		}

		content += "Content for " + route.Path + ": " + markdown + "\n"
	}

	return nil, FetchCategoryMarkdownOutput{MarkdownContent: content}, nil
}

func FetchCategories(ctx context.Context, req *mcp.CallToolRequest, input FetchCategoriesInput) (*mcp.CallToolResult, FetchCategoriesOutput, error) {
	var categories []string

	for _, value := range utils.Routes {
		if !strings.Contains(value.Path, "/") {
			continue
		}

		// strip the filename at the end of the path
		pathItems := strings.Split(value.Path, "/")

		pathItems = slices.Delete(pathItems, len(pathItems)-1, len(pathItems))
		path := strings.Join(pathItems, "/")

		if !slices.Contains(categories, path) {
			categories = append(categories, path)
		}
	}

	return nil, FetchCategoriesOutput{Categories: utils.FormatCategories(categories)}, nil
}
