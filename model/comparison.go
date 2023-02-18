package model

func (f *FileTree) ContainsFile(fileToCompare *File) (bool, *File) {
	var contains = false
	var file *File = nil

	for _, file := range f.Tree.Files {
		if isEqualFile(fileToCompare, file) {
			return true, file
		}
	}

	for _, folder := range f.Tree.Folders {
		TraverseFiles(folder.Files, func(x *File) {
			if isEqualFile(fileToCompare, x) {
				contains = true
				file = x
				return
			}
		})

		for _, nextFolder := range folder.Folders {
			TraverseFilesAndFolders(nextFolder, func(x *File) {
				if isEqualFile(fileToCompare, x) {
					contains = true
					file = x
					return
				}
			}, nil)
		}
	}

	return contains, file
}

func (f *Folder) AddFile(fileToAdd *File) {
	f.Files = append(f.Files, fileToAdd)
}

func (f *Folder) AddFolder(newFolderName string) *Folder {
	var newFolder = &Folder{Name: newFolderName}
	f.Folders = append(f.Folders, newFolder)
	return newFolder
}

func isEqualFile(a, b *File) bool {
	return a != nil && b != nil &&
		a.Name == b.Name &&
		a.Size == b.Size &&
		a.LastModified.UTC() == b.LastModified.UTC()
}

func isEqualFolder(a, b *Folder) bool {
	return a != nil && b != nil &&
		a.Name == b.Name &&
		len(a.Files) == len(b.Files) &&
		len(a.Folders) == len(b.Folders)
}
