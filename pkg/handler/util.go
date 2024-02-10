package handler

import (
    "fmt"
    "errors"
    "os"
    "path/filepath"
)

func IdStealer(handlers []*StealerHandler, dir string) (string, error) {
    for _, h := range handlers {
        fmt.Printf("[*] Testing '%s'\n", h.HandlerName)
        if h.CheckFunction(dir) {
            return h.StealerName, nil
        }
        fmt.Println("[-] No match")
    }
    return "", errors.New("Unknown stealer family")
}

func CheckFiles(dir string, paths []string) bool {
    for _, p := range paths {
        fullPath := filepath.Join(dir, p)
        if _, err := os.Stat(fullPath); os.IsNotExist(err) {
            return false
        }
        fmt.Printf("[+] '%s' exists\n", p)
    }
    return true
}
