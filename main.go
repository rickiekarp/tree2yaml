package main

import (
	"flag"
	"os"

	"git.rickiekarp.net/rickie/tree2yaml/generator"
	"git.rickiekarp.net/rickie/tree2yaml/loader"
	"git.rickiekarp.net/rickie/tree2yaml/printer"
)

func main() {

	var flagGenerate = flag.Bool("generate", true, "generates a filelist")
	var flagLoadList = flag.Bool("load", false, "loads an existing filelist")
	var flagHelp = flag.Bool("help", false, "prints all available options")

	flag.Parse()

	if *flagHelp {
		printer.PrintUsage()
	}

	if flag.Args()[0] == "" {
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
