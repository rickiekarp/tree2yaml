package model

func (f *Folder) AddFile(fileToAdd *File) {
	f.Files = append(f.Files, fileToAdd)
}

func (f *Folder) AddFolder(newFolderName string) *Folder {
	var newFolder = &Folder{Name: newFolderName}
	f.Folders = append(f.Folders, newFolder)
	return newFolder
}
