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
    RequiredFiles: []string{
        "/UserInformation.txt",
    },
}


var rootDirRegex = regexp.MustCompile(`(?m)[A-Z]{2}\[?[A-Z\d]{32}\]?[\s_]\[?\d{4}_\d{2}_\d{2}T\d{2}_\d{2}_\d{2}(_\d{6,10}|\.\d{7,10}(\]|-\d{2}_\d{2}\]))`)

func check(sample_path string) bool {
    val := rootDirRegex.MatchString(sample_path)
    if !val {
        return false
    }
    fmt.Println("[+] Root directory named matched")
    fmt.Println(val)
    return val
}
