package model

type fileIterationFn func(*File)
type folderIterationFn func(*Folder)

func TraverseFiles(files []*File, onFileIteration fileIterationFn) {
	if onFileIteration == nil {
		return
	}
	for _, file := range files {
		onFileIteration(file)
	}
}

func TraverseFilesAndFolders(folders *Folder, onFileIteration fileIterationFn, onFolderIteration folderIterationFn) {
	TraverseFiles(folders.Files, onFileIteration)
	TraverseFolder(folders.Folders, onFolderIteration)

	for _, folder := range folders.Folders {
		TraverseFilesAndFolders(folder, onFileIteration, onFolderIteration)
	}
}

func TraverseFolder(folders []*Folder, onFolderIteration folderIterationFn) {
	if onFolderIteration == nil {
		return
	}
	for _, folder := range folders {
		onFolderIteration(folder)
	}
}
