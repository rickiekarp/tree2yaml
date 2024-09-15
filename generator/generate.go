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

	"git.rickiekarp.net/rickie/tree2yaml/eventsender"
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
var flagGenerateArchive = flag.Bool("enableArchive", false, "generates archive of the generated filelist")
var flagLoadFromFile = flag.Bool("generateMetadataFromFile", false, "load a file list file")

func Generate(filePath string) {

	var tree *model.FileTree = nil
	if *flagLoadFromFile {
		tree = model.LoadFilelist(filePath)
	} else {
		tree = buildTree(filePath, *flagFileHashMethod)
	}

	if len(*flagOutFile) > 0 {
		writeFiletreeToFile(tree, *flagOutFile)

		if *flagGenerateMetadata {
			GenerateAdditionalData(tree, Metadata)
		}

		if *flagGenerateArchive {
			GenerateAdditionalData(tree, Archive)
		}

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

func GenerateAdditionalData(currentFileTree *model.FileTree, generationType GenerationType) {
	var dataFile = *flagOutFile + "." + generationType.String()
	var outFiletree *model.FileTree = nil

	switch generationType {
	case Archive:
		metaDataFile := *flagOutFile + "." + Metadata.String()
		if extensions.FileExists(metaDataFile) {
			metadataFiletree := model.LoadFilelist(metaDataFile)
			var archiveMap map[uint64]model.FileArchive = nil
			if extensions.FileExists(dataFile) {
				fileArchiveMap := model.LoadFileArchive(dataFile)
				archiveMap = updateArchive(metadataFiletree, fileArchiveMap)
			} else {
				archiveMap = createArchive(currentFileTree)
			}
			writeFileArchiveToFile(archiveMap, dataFile)
		} else {
			archiveMap := createArchive(currentFileTree)
			writeFileArchiveToFile(archiveMap, dataFile)
		}
	case Metadata:
		if extensions.FileExists(dataFile) {
			metaFileTree := model.LoadFilelist(dataFile)
			switch generationType {
			case Metadata:
				outFiletree = updateMetadata(currentFileTree, metaFileTree)
				writeFiletreeToFile(outFiletree, dataFile)
			}
		} else {
			outFiletree = addMetadataToFiles(currentFileTree)
			writeFiletreeToFile(outFiletree, dataFile)
		}
	}
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

			pathWithoutRoot := strings.Replace(path.Dir(filePath), rootDir, "", 1)
			pathWithoutRoot = strings.TrimPrefix(pathWithoutRoot, "/")

			var fileToAdd = &model.File{
				Path:         pathWithoutRoot,
				Name:         path.Base(filePath),
				Size:         info.Size(),
				LastModified: info.ModTime(),
			}

			eventsender.SendEventForFile(*fileToAdd)

			// deprecated
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

func writeFileArchiveToFile(archive map[uint64]model.FileArchive, outFile string) {
	data, err := yaml.Marshal(archive)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(outFile, data, 0644)
	if err != nil {
		os.Exit(1)
	}
}
