package main

import (
	"flag"
	"fmt"
	"os"

	"git.rickiekarp.net/rickie/goutilkit"
	"git.rickiekarp.net/rickie/tree2yaml/generator"
	"git.rickiekarp.net/rickie/tree2yaml/loader"
)

func main() {

	var flagGenerate = flag.Bool("generate", true, "generates a filelist")
	var flagLoadList = flag.Bool("load", false, "loads an existing filelist")
	var flagHelp = flag.Bool("help", false, "prints all available options")
	var flagVersion = flag.Bool("v", false, "prints version")

	flag.Parse()

	if *flagHelp {
		goutilkit.PrintUsageAndExit()
	}

	if *flagVersion {
		fmt.Println(generator.Version)
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		fmt.Println("No arguments provided!")
		goutilkit.PrintUsage()
		os.Exit(1)
	}
	filePath := flag.Args()[0]

	// if flagLoadList was set to true, make sure we skip filelist generation
	if *flagLoadList && *flagGenerate {
		*flagGenerate = false
	}

	if *flagGenerate {
		generator.Generate(filePath)
	}

	if *flagLoadList {
		loader.Load(filePath)
	}
}
