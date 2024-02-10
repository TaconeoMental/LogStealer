package main

import (
    "fmt"
    "madoka.pink/logstealer/pkg/handler"
)

var METADATA = handler.StealerHandler{
    HandlerName: "Redline Handler",
    StealerName: "Redline Stealer",
    CheckFunction: check,
}

func check(sample_path string) bool {
    fmt.Println("I'm in redline.go!")
    return false
}
