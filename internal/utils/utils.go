package utils

import (
	"fmt"
	"strings"
)

func ExtractHash(url string) (hash string) {
	split := strings.Split(url, "/")
	return split[len(split)-1]
}

func FormatCategories(categories []string) (table string) {
	markdownTable := fmt.Sprintf("|%s|\n", "category") + fmt.Sprintf("|%s|\n", "--------")

	for _, category := range categories {
		markdownTable += fmt.Sprintf("|%s|\n", category)
	}

	return markdownTable
}
