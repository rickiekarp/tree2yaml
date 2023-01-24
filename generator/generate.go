package generator

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"

	"git.rickiekarp.net/rickie/tree2yaml/hash"
	"git.rickiekarp.net/rickie/tree2yaml/model"
	"git.rickiekarp.net/rickie/tree2yaml/sorting"
)

var Version = "development" // Version set during go build using ldflags

func BuildTree(rootDir string, flagCalcMd5 *bool) *model.FileTree {
	rootDir = path.Clean(rootDir)

	var filetree *model.FileTree = &model.FileTree{}
	var tree *model.Folder
	var nodes = map[string]interface{}{}
	var walkFunc filepath.WalkFunc = func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			nodes[p] = &model.Folder{
				Name:    path.Base(p),
				Files:   []*model.File{},
				Folders: []*model.Folder{},
			}
		} else {
			var md5 = ""
			if *flagCalcMd5 {
				md5 = hash.CalcMd5(p)
			}

			filetree.Size += info.Size()

			nodes[p] = &model.File{
				Name:         path.Base(p),
				Size:         info.Size(),
				LastModified: info.ModTime(),
				Md5:          md5,
			}
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
