package generator

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"git.rickiekarp.net/rickie/tree2yaml/extensions"
	"git.rickiekarp.net/rickie/tree2yaml/hash"
	"git.rickiekarp.net/rickie/tree2yaml/model"
	"git.rickiekarp.net/rickie/tree2yaml/sorting"
	"gopkg.in/yaml.v2"
)

var Version = "development" // Version set during go build using ldflags

var flagFileHashMethod = flag.String("hash", "", "calculate hash sum of each file (crc32, crc64, md5)")

var flagOutFile = flag.String("outfile", "", "path of the output file")
var flagGenerateMetadata = flag.Bool("enableMetadata", false, "generates metadata of the generated filelist")

func Generate(filePath string) {

	tree := buildTree(filePath, *flagFileHashMethod)

	data, err := yaml.Marshal(&tree)
	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
		os.Exit(1)
	}

	if len(*flagOutFile) > 0 {
		err = os.WriteFile(*flagOutFile, data, 0644)
		if err != nil {
			os.Exit(1)
		}

		if *flagGenerateMetadata {
			var metadataFile = *flagOutFile + ".meta"
			if extensions.FileExists(metadataFile) {
				//TODO
			} else {
				//TODO
				err = os.WriteFile(metadataFile, data, 0644)
				if err != nil {
					os.Exit(1)
				}
			}
		}

	} else {
		fmt.Println(string(data))
	}

	os.Exit(0)
}

func buildTree(rootDir string, flagFileHashMethod string) *model.FileTree {
	rootDir = path.Clean(rootDir)

	var filetree *model.FileTree = &model.FileTree{}
	var tree *model.Folder
	var nodes = map[string]interface{}{}
	var walkFunc filepath.WalkFunc = func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			nodes[filePath] = &model.Folder{
				Name:    path.Base(filePath),
				Files:   []*model.File{},
				Folders: []*model.Folder{},
			}
		} else {

			filetree.Size += info.Size()

			var fileToAdd = &model.File{
				Name:         path.Base(filePath),
				Size:         info.Size(),
				LastModified: info.ModTime(),
			}

			if len(flagFileHashMethod) > 0 {
				hashMethods := strings.Split(flagFileHashMethod, ",")
				for _, hashMethod := range hashMethods {
					calcHashByMethod(hashMethod, fileToAdd, filePath)
				}
			}

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

func calcHashByMethod(hashMethod string, file *model.File, filePath string) {
	switch hashMethod {
	case "crc32":
		crc32 := hash.CalcCrc32(filePath)
		file.Crc32 = crc32
	case "crc64":
		crc64 := hash.CalcCrc64(filePath)
		file.Crc64 = crc64
	case "md5":
		md5 := hash.CalcMd5(filePath)
		file.Md5 = md5
	}
}
