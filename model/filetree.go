package model

import (
	"fmt"
	"time"

	"gopkg.in/yaml.v2"
)

type FileTree struct {
	RootDir       string `yaml:"directory"`
	ParserVersion string `yaml:"parserVersion"`
	Size          int64  `yaml:"size"`
	Tree          *Folder
}

type Folder struct {
	Name    string
	Files   []*File   `yaml:"files,omitempty"`
	Folders []*Folder `yaml:"folders,omitempty"`
}

func (f *Folder) String() string {
	j, _ := yaml.Marshal(f)
	return string(j)
}

type File struct {
	Name         string
	Size         int64
	LastModified time.Time
	Md5          string `yaml:"md5,omitempty"`
}

func (f *File) String() string {
	return fmt.Sprintf("%s", f.Name)
}
