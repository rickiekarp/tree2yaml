package generator

import (
	"git.rickiekarp.net/rickie/tree2yaml/model"
)

func createArchive(filetree *model.FileTree) map[uint64]model.FileArchive {
	fileArchive := make(map[uint64]model.FileArchive)
	return updateArchive(filetree, fileArchive)
}

func updateArchive(filetree *model.FileTree, archive map[uint64]model.FileArchive) map[uint64]model.FileArchive {
	model.TraverseFilesAndFolders(filetree.Tree, func(x *model.File) {
		archive[x.Hash()] = model.FileArchive{
			Name:     x.Name,
			Revision: x.Metadata.Revision}

	}, nil)

	return archive
}
