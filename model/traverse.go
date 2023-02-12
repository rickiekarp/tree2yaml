package model

type fileIterationFn func(*File)

func TraverseFiles(files []*File, onFileIteration fileIterationFn) {
	for _, file := range files {
		onFileIteration(file)
	}
}

func TraverseFolders(folders *Folder, onFileIteration fileIterationFn) {
	TraverseFiles(folders.Files, onFileIteration)

	for _, folder := range folders.Folders {
		TraverseFolders(folder, onFileIteration)
	}
}
