package player

import (
	"errors"
	"path/filepath"
	"slices"
	"strings"
)

func isFileValid(path string) bool {
	fileTypes := []string{".mp4", ".mkv", ".avi", ".mov", ".wmv", ".flv", ".webm", ".flv", ".webm", ".m4v"}
	ext := strings.ToLower(filepath.Ext(path))
	return slices.Contains(fileTypes, ext)
}

func normalizePath(path string) (string, error) {
	if path == ("") {
		return "", errors.New("empty string as path")
	}
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	return filepath.Clean(absolutePath), nil
}
