package main

import (
	"flag"
    "fmt"
    "os"
    "log"
    "io/ioutil"
    "strings"
    "path/filepath"
    "madoka.pink/logstealer/pkg/stealerlog"
)

func listSharedLibs(dir string) ([]string, error) {
    objects, err := ioutil.ReadDir(dir)
    if err != nil {
        return nil, err
    }

    var sharedLibs []string
    for _, o := range objects {
        if o.IsDir() || !strings.HasSuffix(o.Name(), ".so") {
            continue
        }
        fullPath := filepath.Join(dir, o.Name())
        sharedLibs = append(sharedLibs, fullPath)
    }
    return sharedLibs, nil
}

func loadHandlers(dir string) ([]*stealerlog.StealerLogHandler, error) {
    sharedLibs, err := listSharedLibs(dir)
    if err != nil {
        return nil, err
    }

    var loghandlers []*stealerlog.StealerLogHandler
    for _, lib := range sharedLibs {
        h, err := stealerlog.Load(lib)
        if err != nil {
            fmt.Printf("[-] Could not load library @ '%s'\n", lib)
            continue
        }
        loghandlers = append(loghandlers, h)
    }
    return loghandlers, nil
}

func testHandler(shandler *stealerlog.StealerLogHandler, path string) bool {
    return shandler.CheckFunction(path)
}

func main() {
	sampleDir := flag.String("sample-dir", "", "InfoStealer sample directory")
	libsDir := flag.String("libs-dir", "", "Compiled libraries directory")

	flag.Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage: %s [Options]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "    -%v,\t%v\n", f.Name, f.Usage)
		})
	}
	flag.Parse()

	if *sampleDir == "" || *libsDir == "" {
		flag.Usage()
		return
    }

    stealerHandlers, err := loadHandlers(*libsDir)
	if err != nil {
		log.Fatal(err.Error())
	    return
    }
		
	fmt.Printf("[+] %d stealers loaded\n", len(stealerHandlers))
    fmt.Println(stealerHandlers)

    stealer, err := stealerlog.IdStealer(stealerHandlers, *sampleDir)
    if err != nil {
        log.Fatal(err.Error())
    }
    fmt.Printf("[+] '%s' detected!\n", stealer)
}
	
