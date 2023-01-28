package extractor

import (
	"git.rickiekarp.net/rickie/tree2yaml/model"
)

func FindFolder(fileTreeRootFolder *model.Folder, searchPathSlice []string) *model.Folder {

	// return file tree root directory if no sub directories have been provided
	if len(searchPathSlice) == 1 {
		return fileTreeRootFolder
	}

	var foundDir *model.Folder
	var rootFolders = fileTreeRootFolder.Folders

	for _, path := range searchPathSlice {

		for _, folder := range rootFolders {

			if folder.Name == path && searchPathSlice[len(searchPathSlice)-1] == folder.Name {
				return folder
			}

			foundDir = findSubDirectory(folder.Folders, path)

			if foundDir != nil {
				rootFolders = folder.Folders
			}
		}
	}

	return foundDir
}

func findSubDirectory(folderToSearch []*model.Folder, path string) *model.Folder {
	for _, folder := range folderToSearch {
		if folder.Name == path {
			return folder
		}
	}

	return nil
}
