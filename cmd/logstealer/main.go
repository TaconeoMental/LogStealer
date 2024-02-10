package main

import (
	"flag"
    "fmt"
    "os"
    "log"
    "io/ioutil"
    "strings"
    "path/filepath"
    "madoka.pink/logstealer/pkg/handler"
)

/*
func GuessStealerFamily(sample_path string, config_path string) (string, error) {
    // Check if both paths exist
    if _, err := os.Stat(sample_path); os.IsNotExist(err) {
        return "", errors.New("Sample file path does not exist")
    }

    if _, err := os.Stat(config_path); os.IsNotExist(err) {
        return "", errors.New("Config file path does not exist")
    }

    plugins := plugin.LoadPlugins(config_path)
    fmt.Printf("[+] %d plugins loaded\n", len(plugins))

    for _, plugin := range plugins {
        check_func, err := plugin.Lookup(checkSymbolName)
        if err != nil {
            log.Fatal(err)
        }

        if check_func.(checkFuncSignature)(sample_path) {

        }
    }

    //////////

    //for _, plugin := range plugins {
    //}

    return "sfkjh", nil
}
*/

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
        full_path := filepath.Join(dir, o.Name())
        sharedLibs = append(sharedLibs, full_path)
    }
    return sharedLibs, nil
}

func loadHandlers(dir string) ([]*handler.StealerHandler, error) {
    sharedLibs, err := listSharedLibs(dir)
    if err != nil {
        return nil, err
    }

    var handlers []*handler.StealerHandler
    for _, lib := range sharedLibs {
        h, err := handler.Load(lib)
        if err != nil {
            fmt.Printf("[-] Could not load library @ '%s'\n", lib)
            continue
        }
        handlers = append(handlers, h)
    }
    return handlers, nil
}

func testHandler(shandler *handler.StealerHandler, path string) bool {
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

    stealer, err := handler.IdStealer(stealerHandlers, *sampleDir)
    if err != nil {
        log.Fatal(err.Error())
    }
    fmt.Printf("[+] '%s' it is!\n", stealer)
}
	
