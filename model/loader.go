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
