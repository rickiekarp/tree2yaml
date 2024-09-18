package printer

import (
	"fmt"
	"sort"
	"strings"

	"git.rickiekarp.net/rickie/tree2yaml/eventsender"
	"git.rickiekarp.net/rickie/tree2yaml/model"
	"gopkg.in/yaml.v2"
)

func PrintFilelist(filelist *model.FileTree) {
	out, err := yaml.Marshal(filelist)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}

func PrintFileListWithOccurrences(fileOccurrenceMap map[int][]model.File) {
	keys := make([]int, 0, len(fileOccurrenceMap))
	for k := range fileOccurrenceMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, k := range keys {
		resultText := fmt.Sprintf("Results (Probability: %d%%)", k)
		fmt.Println(resultText)

		fileSlice := fileOccurrenceMap[k]

		for _, i := range fileSlice {
			fmt.Println(i.Name)
		}
		fmt.Println()
	}
}

func PrintArchive(archive map[uint64]model.FileArchive) {
	out, err := yaml.Marshal(archive)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}

func PrintAllFilesAndSendFileEvents(filelist *model.FileTree) {
	path := filelist.RootDir

	for _, file := range filelist.Tree.Files {
		file.Path = path
		fmt.Println(file.Name)
	}

	var rootFolders = filelist.Tree.Folders
	for _, folder := range rootFolders {
		subPath := folder.Name + "/"
		traverseFiles(folder.Files, subPath)

		for _, nextFolder := range folder.Folders {
			subPath = folder.Name + "/" + nextFolder.Name + "/"
			traverseFolders(nextFolder, subPath)
		}
	}
}

func traverseFiles(files []*model.File, path string) {
	for _, file := range files {
		file.Path += path
		file.Path = strings.TrimSuffix(file.Path, "/")
		fmt.Println(file.Name, " ", file.Path)
		eventsender.SendEventForFile(*file)
	}
}

func traverseFolders(folders *model.Folder, path string) {
	traverseFiles(folders.Files, path)

	for _, folder := range folders.Folders {
		subPath := path + folder.Name + "/"
		traverseFolders(folder, subPath)
	}
}
