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
