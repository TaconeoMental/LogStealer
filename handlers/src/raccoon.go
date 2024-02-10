package main

import (
    "fmt"
    "madoka.pink/logstealer/pkg/handler"
)

var METADATA = handler.StealerHandler{
    HandlerName: "Raccoon Handler",
    StealerName: "RaccoonStealer",
    CheckFunction: check,
}

// Check what happens if symbol is not exported
func check(sample_path string) bool {
    fmt.Println("I'm in racoon.go!")
    return false
}
