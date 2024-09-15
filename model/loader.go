package model

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadFilelist(filePath string) *FileTree {
	var fileTree *FileTree
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &fileTree)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return fileTree
}

func LoadFilelistFromString(fileContent string) *FileTree {
	var fileTree *FileTree
	byteSlice := []byte(fileContent)

	err := yaml.Unmarshal(byteSlice, &fileTree)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return fileTree
}

func LoadFileArchive(filePath string) map[uint64]FileArchive {
	var archive map[uint64]FileArchive
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &archive)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return archive
}
