package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchSingleMarkdown(url string) (string, error) {
	success, content, err := SearchCache(ExtractHash(url))

	if success == true {
		return content, nil
	}

	if err != nil { // SearchCache erroring doesn't really matter, it can be ignored and the content can be fetched fresh
		fmt.Println(err)
	}

	resp, err1 := http.Get(url)

	if err1 != nil {
		return "", err1
	}

	bodyBytes, err2 := io.ReadAll(resp.Body)

	if err2 != nil {
		return "", err2
	}

	b64Body := string(bodyBytes)
	var body map[string]interface{}

	err3 := json.Unmarshal([]byte(b64Body), &body)

	if err3 != nil {
		return "", err3
	}

	rawVal, ok := body["content"]

	if !ok {
		return "", &ReturnError{Message: "Invalid return format from " + url}
	}

	val := rawVal.(string)

	decode, err4 := base64.StdEncoding.DecodeString(val)

	if err4 != nil {
		return "", err4
	}

	markdownContent := string(decode)

	AddCache(ExtractHash(url), markdownContent)

	return markdownContent, nil
}
