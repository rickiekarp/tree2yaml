package loader

import (
	"io/ioutil"
	"log"

	"git.rickiekarp.net/rickie/tree2yaml/model"
	"gopkg.in/yaml.v2"
)

func LoadFilelist(filePath string) *model.FileTree {
	var fileTree *model.FileTree
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &fileTree)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return fileTree
}

func loadFilelistFromString(fileContent string) *model.FileTree {
	var fileTree *model.FileTree
	byteSlice := []byte(fileContent)

	err := yaml.Unmarshal(byteSlice, &fileTree)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return fileTree
}

func LoadFileArchive(filePath string) map[uint64]model.FileArchive {
	var archive map[uint64]model.FileArchive
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &archive)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return archive
}
