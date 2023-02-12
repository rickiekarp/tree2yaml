package model

import (
	"time"
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

type File struct {
	Name         string
	Size         int64
	LastModified time.Time
	Crc32        uint32       `yaml:"crc32,omitempty"`
	Crc64        uint64       `yaml:"crc64,omitempty"`
	Md5          string       `yaml:"md5,omitempty"`
	Metadata     FileMetadata `yaml:"metadata,omitempty"`
}

type FileMetadata struct {
	Revision int64
}
