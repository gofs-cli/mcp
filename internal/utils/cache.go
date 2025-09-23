package utils

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var TTLDays int = 1
var cacheTTL int = TTLDays * 24 * 60 * 60 * 1000
var cacheFileExt string = ".txt"

func SearchCache(hash string) (exists bool, content string, err error) {
	tempPath := filepath.Join(os.TempDir(), "gofs-mcp")

	filePath := filepath.Join(tempPath, hash+cacheFileExt)

	if _, err1 := os.Stat(filePath); err1 != nil {
		if os.IsNotExist(err) {
			return false, "", nil
		} else {
			return false, "", &ReturnError{Message: "Checking existence of temp file called " + hash + cacheFileExt + " failed."}
		}
	} else {
		data, err3 := os.ReadFile(filePath)

		if err3 != nil {
			return false, "", &ReturnError{Message: "Could not read data from found matching file"}
		}

		fileContent := string(data)
		timestampStr, content, found := strings.Cut(fileContent, "\n")

		if !found {
			return false, "", &ReturnError{Message: "Invalid format in found cache file"}
		}

		cachedTimestamp, err4 := strconv.Atoi(strings.TrimSpace(timestampStr))

		if err4 != nil {
			return false, "", &ReturnError{Message: "Could not read timestamp from cache file"}
		}

		currentTimestamp := time.Now().UnixMilli()

		if int64(cachedTimestamp+cacheTTL) < currentTimestamp {
			err5 := os.Remove(filePath)

			if err5 != nil {
				return false, "", &ReturnError{Message: "Could not delete cache file that exceeded TTL"}
			}

			return false, "", nil
		}
		return true, content, nil
	}
}

func AddCache(hash string, content string) (success bool, err error) {
	// make sure gofs-mcp folder exists in the temp dir
	tempPath := filepath.Join(os.TempDir(), "gofs-mcp")

	err1 := os.MkdirAll(tempPath, os.ModePerm) // automatically handles the directory already existing

	if err1 != nil {
		return false, &ReturnError{Message: "Could not make gofs-mcp folder in temp dir"}
	}

	timestamp := time.Now().UnixMilli()
	updatedContent := strconv.FormatUint(uint64(timestamp), 10) + "\n" + content

	filePath := filepath.Join(tempPath, hash+cacheFileExt)

	err2 := os.WriteFile(filePath, []byte(updatedContent), 0644)
	if err2 != nil {
		return false, &ReturnError{Message: "Could not write to cache file at " + filePath}
	}

	return true, nil
}
