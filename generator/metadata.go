package generator

import (
	"fmt"
	"os"

	"git.rickiekarp.net/rickie/tree2yaml/extensions"
	"git.rickiekarp.net/rickie/tree2yaml/loader"
	"git.rickiekarp.net/rickie/tree2yaml/model"
	"gopkg.in/yaml.v2"
)

func GenerateMetadata(filetree *model.FileTree) {
	var metadataFile = *flagOutFile + ".meta"

	if extensions.FileExists(metadataFile) {
		metadataFiletree := loader.LoadFilelist(metadataFile)
		newMetadata := updateMetadata(filetree, metadataFiletree)
		writeMetadataToFile(newMetadata, metadataFile)
	} else {
		newFiletree := addMetadataToFiles(filetree)
		writeMetadataToFile(newFiletree, metadataFile)
	}
}

func addMetadataToFiles(filetree *model.FileTree) *model.FileTree {

	model.TraverseFiles(filetree.Tree.Files, func(x *model.File) {
		x.Metadata = model.FileMetadata{
			Revision: 1,
		}
	})

	for _, folder := range filetree.Tree.Folders {
		model.TraverseFiles(folder.Files, func(x *model.File) {
			x.Metadata = model.FileMetadata{
				Revision: 1,
			}
		})

		for _, nextFolder := range folder.Folders {
			model.TraverseFolders(nextFolder, func(x *model.File) {
				x.Metadata = model.FileMetadata{
					Revision: 1,
				}
			})
		}
	}

	return filetree
}

func updateMetadata(filetree *model.FileTree, metadata *model.FileTree) *model.FileTree {

	model.TraverseFiles(filetree.Tree.Files, func(x *model.File) {
		// if a metadata file was found in the file tree, increment the revision
		contains, file := metadata.ContainsFile(x)
		if contains {
			x.Metadata.Revision = file.Metadata.Revision + 1
		} else {
			x.Metadata = model.FileMetadata{
				Revision: 1,
			}
		}
	})

	for _, folder := range filetree.Tree.Folders {

		model.TraverseFiles(folder.Files, func(x *model.File) {
			// if a metadata file was found in the file tree, increment the revision
			contains, file := metadata.ContainsFile(x)
			if contains {
				x.Metadata.Revision = file.Metadata.Revision + 1
			} else {
				x.Metadata = model.FileMetadata{
					Revision: 1,
				}
			}
		})

		for _, nextFolder := range folder.Folders {
			model.TraverseFolders(nextFolder, func(x *model.File) {
				// if a metadata file was found in the file tree, increment the revision
				contains, file := metadata.ContainsFile(x)
				if contains {
					x.Metadata.Revision = file.Metadata.Revision + 1
				} else {
					x.Metadata = model.FileMetadata{
						Revision: 1,
					}
				}
			})
		}
	}

	return filetree
}

func writeMetadataToFile(filetree *model.FileTree, metadataFile string) {
	data, err := yaml.Marshal(&filetree)
	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
		os.Exit(1)
	}

	err = os.WriteFile(metadataFile, data, 0644)
	if err != nil {
		os.Exit(1)
	}
}
