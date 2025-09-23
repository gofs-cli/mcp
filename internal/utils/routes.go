package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

var Routes []RouteData

type RouteData struct {
	Path string
	URL  string
}

func FormatRoutes(data []RouteData) string {
	formatted := ""

	for _, value := range data {
		temp := "Name: " + value.Path + " which can be accessed at URL: " + value.URL

		formatted += temp + ", "
	}

	return formatted
}

func GetRoutes() ([]RouteData, error) {
	tree_url := "https://api.github.com/repos/gofs-cli/web/git/trees/main?recursive=1"
	resp, err1 := http.Get(tree_url)

	if err1 != nil {
		return nil, err1
	}

	bodyBytes, err2 := io.ReadAll(resp.Body)

	if err2 != nil {
		return nil, err2
	}

	var bodyJson map[string]interface{}

	err3 := json.Unmarshal(bodyBytes, &bodyJson)

	if err3 != nil {
		return nil, err3
	}

	var finalData []RouteData

	for _, value := range bodyJson["tree"].([]interface{}) {
		assertedValue, ok := value.(map[string]interface{})

		if ok { // if the type assertion fails somehow then skip the value
			path := assertedValue["path"].(string)

			if !strings.HasPrefix(path, "docs/") { // only use files in the docs folder
				continue
			}

			if assertedValue["type"].(string) == "tree" { // if the entry is a folder not a file then skip it
				continue
			}

			if strings.Contains(path, "_category_") { // category pages do not have any actual information in them so they can be skipped
				continue
			}

			newRouteData := RouteData{
				Path: assertedValue["path"].(string),
				URL:  assertedValue["url"].(string),
			}

			finalData = append(finalData, newRouteData)
		}
	}
	return finalData, nil
}
