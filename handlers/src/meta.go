package main

// DumpHandler

import (
    "fmt"
    "madoka.pink/logstealer/pkg/stealerlog"
)

var HANDLER = stealerlog.StealerLogHandler{
    HandlerName: "Meta Handler",
    StealerName: "MetaStealer",
    CheckFunction: check,
}

var globalSignature = `(\( M \| E \| T \| A \)|https://t\.me/metastealer_bot)`
var rootDirSignature = `(?m)[A-Z]{2}\[?[A-Z\d]{32}\]?[\s_]\[?\d{4}_\d{2}_\d{2}T\d{2}_\d{2}_\d{2}(_\d{6,10}|\.\d{7,10}(\]|-\d{2}_\d{2}\]))`

var logRules = []stealerlog.Rule{
    &stealerlog.File{
        Path: "/ProcessList.txt",
        Signature: stealerlog.NewSignature(
            globalSignature,
            `={15}`,
            `ID:\s\d+,\sName:\s.+,\sCommandLine:`,
        ),
    },
    &stealerlog.File{
        Path: "/DomainDetects.txt",
    },
    &stealerlog.File{
        Path: "/InstalledSoftware.txt",
        Signature: stealerlog.NewSignature(globalSignature),
    },
    &stealerlog.File{
        Path: "/UserInformation.txt",
        Signature: stealerlog.NewSignature(
            globalSignature,
            `Build ID:\s.+`,
            `Process Elevation:\s(False|True)`,
        ),
        Extract: xUserInformation,
    },
    &stealerlog.File{
        Path: "/FTP/Credentials.txt",
        Signature: stealerlog.NewSignature(globalSignature),
        Extract: xFTPCredentials,
        Optional: true,
    },
    //stealerlog.Directory{
    //&stealerlog.File{
        //Path: "/Cookies/",
        //Extract: xCookiesDir,
    //},
}

func xUserInformation(path string) {}
func xFTPCredentials(path string) {}
func xCookiesDir(path string) {}

func check(samplePath string) bool {
    if !stealerlog.CheckRules(samplePath, logRules) {
        return false
    }
    fmt.Println("[+] All rules passed")

    return true
}
