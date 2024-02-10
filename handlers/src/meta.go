package main

import (
    "fmt"
    "regexp"
    "madoka.pink/logstealer/pkg/handler"
)

var METADATA = handler.StealerHandler{
    HandlerName: "Meta Handler",
    StealerName: "MetaStealer",
    CheckFunction: check,
}

var rootDirRegex = regexp.MustCompile(`(?m)[A-Z]{2}\[?[A-Z\d]{32}\]?[\s_]\[?\d{4}_\d{2}_\d{2}T\d{2}_\d{2}_\d{2}(_\d{6,10}|\.\d{7,10}(\]|-\d{2}_\d{2}\]))`)
var requiredFiles = []string{"/UserInformation.txt"}

func check(samplePath string) bool {
    if !rootDirRegex.MatchString(samplePath) {
        return false
    }
    fmt.Println("[+] Root directory name matches")

    if !handler.CheckFiles(samplePath, requiredFiles) {
        return false
    }
    fmt.Println("[+] Required files exist")

    return true
}
