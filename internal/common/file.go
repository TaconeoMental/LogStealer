package common

import (
    "errors"
    "os"
    "log"
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
            log.Fatal("Cannot get current working directory")
            return "", err
    }
    expandedPath := filepath.Join(cwdPath, path)
    return expandedPath, nil
}
