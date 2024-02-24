package main

import (
	"flag"
	"fmt"
	"os"

	//"madoka.pink/logstealer/pkg/stealerlog"
	//"github.com/k0kubun/pp"
	"madoka.pink/logstealer/internal/stealer"
)

func main() {
	sampleDir := flag.String("sample-dir", "", "InfoStealer sample directory")
	configFilePath := flag.String("config-file", "", "Stealer config file")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [Options]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "    -%v,\t%v\n", f.Name, f.Usage)
		})
	}
	flag.Parse()

	if *sampleDir == "" || *configFilePath == "" {
		flag.Usage()
		return
	}

	stealer, err := stealer.FromConfigFile(*configFilePath)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    stealer.ExtractData(*sampleDir)
}
