package loader

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"git.rickiekarp.net/rickie/tree2yaml/connector/gitconnector"
	"git.rickiekarp.net/rickie/tree2yaml/extractor"
	"git.rickiekarp.net/rickie/tree2yaml/model"
	"git.rickiekarp.net/rickie/tree2yaml/printer"
)

var flagFindFiles = flag.String("findFilesIn", "", "finds files by a given search path, e.g. tree2yaml --load --findFilesIn=foo/bar /foo/bar.yaml")
var flagFindFolders = flag.String("findFoldersIn", "", "finds folders by a given search path, e.g. tree2yaml --load --findFoldersIn=foo/bar /foo/bar.yaml")
var flagMatchFiles = flag.String("find", "", "prints all file names that match the given arguments, grouped by occurrence, e.g. --find=foo,bar")

var flagFilterByDate = flag.String("filterByDate", "", "filters files by given date, e.g. --filterByDate=2022-12-24")
var flagFilterByDateDirection = flag.String("filterByDateDirection", "new", "direction of files to be filtered, e.g. 'old', 'new'")
var flagIgnoreCase = flag.Bool("ignoreCase", false, "ignore case when matching files, can be combined with -find flag")

var flagGitHistory = flag.Bool("git", false, "check git history")
var flagGitLogDepth = flag.Int("git-depth", 3, "git log depth")

func Load(filePath string) {

	if len(*flagFindFiles) > 0 {

		findFiles(filePath)

	} else if len(*flagFindFolders) > 0 {

		findFolders(filePath)

	} else if len(*flagMatchFiles) > 0 {

		if *flagGitHistory {

			var results = make(map[int][]model.File)

			hashes := gitconnector.GetGitLogHashes(*flagGitLogDepth)
			for _, hash := range hashes {
				fileContent, exitCode := gitconnector.ShowFileAtRevision(filePath, hash)
				if exitCode == 0 {
					fileList := loadFilelistFromString(fileContent)
					var result = extractor.MatchOccurrencesInFileTree(fileList, flagMatchFiles, flagIgnoreCase)
					results = DeepCopy(result, results)

				} else {
					fmt.Println("File not found:", filePath)
				}
			}

			printer.PrintFileListWithOccurrences(results)

		} else {

			filelist := loadFilelist(filePath)
			result := extractor.MatchOccurrencesInFileTree(filelist, flagMatchFiles, flagIgnoreCase)
			printer.PrintFileListWithOccurrences(result)

		}

	} else {

		filelist := loadFilelist(filePath)
		printer.PrintFilelist(filelist)

	}
}

func DeepCopy(src, dst map[int][]model.File) map[int][]model.File {

	var tmp = make(map[model.File]int)
	var newMap = make(map[int][]model.File)

	for srcProbability, srcFiles := range src {
		newMap[srcProbability] = []model.File{}
		for _, file := range srcFiles {
			tmp[file] = srcProbability
		}
	}

	for srcProbability, srcFiles := range dst {
		newMap[srcProbability] = []model.File{}
		for _, file := range srcFiles {
			tmp[file] = srcProbability
		}
	}

	for file, prob := range tmp {
		var entry = newMap[prob]
		entry = append(entry, file)
		newMap[prob] = entry
	}

	return newMap
}

func findFiles(filePath string) {
	filelist := loadFilelist(filePath)
	splitDirectorySlice := strings.Split(*flagFindFiles, "/")
	folder := extractor.FindFolder(filelist.Tree, splitDirectorySlice)
	if folder == nil {
		os.Exit(0)
	}

	folder = extractor.FilterByModDate(folder, *flagFilterByDate, *flagFilterByDateDirection)

	for _, file := range folder.Files {
		fmt.Println(file.Name)
	}
}

func findFolders(filePath string) {
	splitDirectorySlice := strings.Split(*flagFindFolders, "/")
	filelist := loadFilelist(filePath)
	folder := extractor.FindFolder(filelist.Tree, splitDirectorySlice)
	if folder == nil {
		os.Exit(0)
	}
	for _, folder := range folder.Folders {
		fmt.Println(folder.Name)
	}
}
