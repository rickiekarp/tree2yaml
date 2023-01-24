package extractor

import "git.rickiekarp.net/rickie/tree2yaml/model"

func FindFolder(slice []*model.Folder, searchPathSlice []string) *model.Folder {

	var foundDir *model.Folder

	for _, path := range searchPathSlice {

		for _, folder := range slice {

			if folder.Name == path && searchPathSlice[len(searchPathSlice)-1] == folder.Name {
				return folder
			}

			foundDir = findSubDirectory(folder.Folders, path)

			if foundDir != nil {
				slice = folder.Folders
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
