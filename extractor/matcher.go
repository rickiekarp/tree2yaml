package extractor

import (
	"strings"

	"git.rickiekarp.net/rickie/tree2yaml/extensions"
	"git.rickiekarp.net/rickie/tree2yaml/model"
)

var matchFileSlice []string
var probabilityMap map[int]int
var fileMap map[int][]model.File
var ignoreCase bool

func MatchOccurrencesInFileTree(filetree *model.FileTree, flagMatchFiles *string, flagIgnoreCase *bool) map[int][]model.File {
	matchFileSlice = strings.Split(*flagMatchFiles, ",")
	probabilityMap = getOccurrenceProbabilityMap(matchFileSlice)
	fileMap = make(map[int][]model.File)
	ignoreCase = *flagIgnoreCase

	var rootFolders = filetree.Tree.Folders
	for _, folder := range rootFolders {

		traverseFiles(folder.Files)

		for _, nextFolder := range folder.Folders {
			traverseFolders(nextFolder)
		}
	}

	return fileMap
}

func traverseFiles(files []*model.File) {
	for _, file := range files {
		if contains, occurrences := containsFileWithOccurrences(file, matchFileSlice); contains {
			fileMap[probabilityMap[occurrences]] = append(fileMap[probabilityMap[occurrences]], *file)
		}
	}
}

func traverseFolders(folders *model.Folder) {
	traverseFiles(folders.Files)

	for _, folder := range folders.Folders {
		traverseFolders(folder)
	}
}

func containsFileWithOccurrences(file *model.File, matchFileSlice []string) (bool, int) {
	occurrences := 0

	if ignoreCase {
		for _, match := range matchFileSlice {
			if extensions.ContainsCaseInsensitive(file.Name, match) {
				occurrences++
			}
		}
	} else {
		for _, match := range matchFileSlice {
			if strings.Contains(file.Name, match) {
				occurrences++
			}
		}
	}

	return occurrences > 0, occurrences
}

func getOccurrenceProbabilityMap(slice []string) map[int]int {
	occurrenceProbabilityMap := make(map[int]int)
	lenSlice := float32(len(slice))

	for i := range slice {
		probability := float32(i+1) / lenSlice * 100
		occurrenceProbabilityMap[i+1] = int(probability)
	}

	return occurrenceProbabilityMap
}
