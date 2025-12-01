package generator

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"git.rickiekarp.net/rickie/tree2yaml/eventsender"
	"git.rickiekarp.net/rickie/tree2yaml/model"
	"git.rickiekarp.net/rickie/tree2yaml/sorting"
	"gopkg.in/yaml.v2"
)

var Version = "development" // Version set during go build using ldflags

var flagOutFile = flag.String("outfile", "", "path of the output file")
var flagLoadFromFile = flag.Bool("generateMetadataFromFile", false, "load a file list file")

func Generate(filePath string) {

	var tree *model.FileTree = nil

	// load from file or build new tree
	if *flagLoadFromFile {
		tree = model.LoadFilelist(filePath)
	} else {
		tree = buildTree(filePath)
	}

	// print or write to file
	if len(*flagOutFile) > 0 {
		writeFiletreeToFile(tree, *flagOutFile)
	} else {
		data, err := marshalFileTree(tree)
		if err != nil {
			fmt.Printf("Error while Marshaling. %v", err)
			os.Exit(1)
		}
		fmt.Println(string(data))
	}

	os.Exit(0)
}

func buildTree(rootDir string) *model.FileTree {
	rootDir = path.Clean(rootDir)

	// generate process ID
	t := time.Now()
	year := t.Year() % 100 // take last 2 digits
	day := t.YearDay()     // 1–365 (or 366)
	hour := t.Hour()
	minute := t.Minute()
	result := fmt.Sprintf("%02d%03d%02d%02d", year, day, hour, minute)
	processId, _ := strconv.ParseInt(result, 10, 64)

	// build file tree
	var filetree *model.FileTree = &model.FileTree{}
	var tree *model.Folder
	var nodes = map[string]any{}
	var walkFunc filepath.WalkFunc = func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			nodes[filePath] = &model.Folder{
				Name:    path.Base(filePath),
				Files:   []*model.File{},
				Folders: []*model.Folder{},
			}
		} else {
			filetree.Size += info.Size()

			finalPath := strings.Replace(path.Dir(filePath), rootDir, "", 1)
			finalPath = strings.TrimPrefix(finalPath, "/")

			// the filePath will be empty if a file is present in the root working directory
			if finalPath == "" {
				finalPath = path.Dir(filePath)
			}

			var fileToAdd = &model.File{
				Path:         finalPath,
				Name:         path.Base(filePath),
				Size:         info.Size(),
				LastModified: info.ModTime(),
			}

			eventsender.SendEventForFile(*fileToAdd, &processId)

			nodes[filePath] = fileToAdd
		}

		return nil
	}

	err := filepath.Walk(rootDir, walkFunc)
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range nodes {
		var parentFolder *model.Folder
		if key == rootDir {
			tree = value.(*model.Folder)
			continue
		} else {
			parentFolder = nodes[path.Dir(key)].(*model.Folder)
		}

		switch v := value.(type) {
		case *model.File:
			parentFolder.Files = append(parentFolder.Files, v)
			sort.Sort(sorting.ByName(parentFolder.Files))
		case *model.Folder:
			parentFolder.Folders = append(parentFolder.Folders, v)
			sort.Sort(sorting.ByFolderName(parentFolder.Folders))
		}
	}

	filetree.RootDir = rootDir
	filetree.ParserVersion = Version
	filetree.Tree = tree

	return filetree
}

func marshalFileTree(filetree *model.FileTree) ([]byte, error) {
	return yaml.Marshal(&filetree)
}

func writeFiletreeToFile(filetree *model.FileTree, outFile string) {
	data, err := marshalFileTree(filetree)
	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
		os.Exit(1)
	}

	err = os.WriteFile(outFile, data, 0644)
	if err != nil {
		os.Exit(1)
	}
}
