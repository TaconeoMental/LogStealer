package handler

import (
    "fmt"
    "errors"
)

func IdStealer(handlers []*StealerHandler, dir string) (string, error) {
    for _, h := range handlers {
        fmt.Printf("[*] Testing '%s'\n", h.HandlerName)
        if h.CheckFunction(dir) {
            return h.StealerName, nil
        }
    }
    return "", errors.New("Unknown stealer family")
}
