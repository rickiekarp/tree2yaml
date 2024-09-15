package generator

import (
	"git.rickiekarp.net/rickie/tree2yaml/model"
)

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
			model.TraverseFilesAndFolders(nextFolder, func(x *model.File) {
				x.Metadata = model.FileMetadata{
					Revision: 1,
				}
			}, nil)
		}
	}

	return filetree
}

func updateMetadata(currentFileTree *model.FileTree, metadata *model.FileTree) *model.FileTree {

	model.TraverseFiles(currentFileTree.Tree.Files, func(x *model.File) {
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

	for _, folder := range currentFileTree.Tree.Folders {

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
			model.TraverseFilesAndFolders(nextFolder, func(x *model.File) {
				// if a metadata file was found in the file tree, increment the revision
				contains, file := metadata.ContainsFile(x)
				if contains {
					x.Metadata.Revision = file.Metadata.Revision + 1
				} else {
					x.Metadata = model.FileMetadata{
						Revision: 1,
					}
				}
			}, nil)
		}
	}

	return currentFileTree
}
