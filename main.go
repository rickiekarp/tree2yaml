package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"git.rickiekarp.net/rickie/tree2yaml/extractor"
	"git.rickiekarp.net/rickie/tree2yaml/generator"
	"git.rickiekarp.net/rickie/tree2yaml/loader"
	"git.rickiekarp.net/rickie/tree2yaml/printer"
	"gopkg.in/yaml.v2"
)

var filePath string

func main() {

	var flagCalcMd5 = flag.Bool("calcMd5", false, "calculate md5 sum of each file")
	var flagGenerate = flag.Bool("generate", true, "generates a filelist")
	var flagLoadList = flag.Bool("load", false, "loads an existing filelist")
	var flagFindFiles = flag.String("findFilesIn", "", "finds files by a given search path, e.g. tree2yaml --load --findFilesIn=foo/bar /foo/bar.yaml")
	var flagFindFolders = flag.String("findFoldersIn", "", "finds folders by a given search path, e.g. tree2yaml --load --findFoldersIn=foo/bar /foo/bar.yaml")
	var flagFilterByDate = flag.String("filterByDate", "", "filters files by given date, e.g. --filterByDate=2022-12-24")
	var flagFilterByDateDirection = flag.String("filterByDateDirection", "new", "direction of files to be filtered, e.g. 'old', 'new'")
	var flagMatchFiles = flag.String("find", "", "prints all file names that match the given arguments, grouped by occurrence, e.g. --find=foo,bar")

	flag.Parse()

	if flag.Args()[0] == "" {
		os.Exit(1)
	}
	filePath = flag.Args()[0]

	// if flagLoadList was set to true, make sure we skip filelist generation
	if *flagLoadList && *flagGenerate {
		*flagGenerate = false
	}

	if *flagGenerate {
		tree := generator.BuildTree(filePath, flagCalcMd5)

		data, err := yaml.Marshal(&tree)

		if err != nil {
			fmt.Printf("Error while Marshaling. %v", err)
			os.Exit(1)
		}

		fmt.Println(string(data))
		os.Exit(0)
	}

	if *flagLoadList {
		filelist := loader.LoadFilelist(filePath)

		if len(*flagFindFiles) > 0 {
			splitDirectorySlice := strings.Split(*flagFindFiles, "/")
			folder := extractor.FindFolder(filelist.Tree, splitDirectorySlice)
			if folder == nil {
				os.Exit(0)
			}

			folder = extractor.FilterByModDate(folder, *flagFilterByDate, *flagFilterByDateDirection)

			for _, file := range folder.Files {
				fmt.Println(file.Name)
			}
		} else if len(*flagFindFolders) > 0 {
			splitDirectorySlice := strings.Split(*flagFindFolders, "/")
			folder := extractor.FindFolder(filelist.Tree, splitDirectorySlice)
			if folder == nil {
				os.Exit(0)
			}
			for _, folder := range folder.Folders {
				fmt.Println(folder.Name)
			}
		} else if len(*flagMatchFiles) > 0 {
			result := extractor.MatchOccurrencesInFileTree(filelist, flagMatchFiles)
			printer.PrintFileListWithOccurrences(result)
		} else {
			printer.PrintFilelist(filelist)
		}
	}
}
