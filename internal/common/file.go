package common

import (
	"errors"
	"os"
	"path/filepath"
)

func FileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		// File doesn't exist
		return false
	}
	return true
}

func ExpandPath(path string) (string, error) {
	cwdPath, err := os.Getwd()
	if err != nil {
		return "", errors.New("Cannot get current working directory")
	}
	expandedPath := filepath.Join(cwdPath, path)
	return expandedPath, nil
}
