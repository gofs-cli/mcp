package utils

import "strings"

func ExtractHash(url string) (hash string) {
	split := strings.Split(url, "/")
	return split[len(split)-1]
}
