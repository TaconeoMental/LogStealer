package stealerlog

import (
    "fmt"
    "errors"
)

func IdStealer(handlers []*StealerLogHandler, dir string) (string, error) {
    for _, h := range handlers {
        fmt.Printf("[*] Testing '%s'\n", h.HandlerName)
        if h.Test(dir) {
            return h.StealerName, nil
        }
        fmt.Println("[-] No match")
    }
    return "", errors.New("Unknown stealer family")
}

