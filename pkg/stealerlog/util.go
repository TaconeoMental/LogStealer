package stealerlog

import (
    "fmt"
    "errors"
//    "os"
//  "path/filepath"
)

func IdStealer(handlers []*StealerLogHandler, dir string) (string, error) {
    for _, h := range handlers {
        fmt.Printf("[*] Testing '%s'\n", h.HandlerName)
        if h.CheckFunction(dir) {
            return h.StealerName, nil
        }
        fmt.Println("[-] No match")
    }
    return "", errors.New("Unknown stealer family")
}

