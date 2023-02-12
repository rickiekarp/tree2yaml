package model

func (f *FileTree) ContainsFile(fileToCompare *File) (bool, *File) {
	var contains = false
	var file *File = nil

	for _, file := range f.Tree.Files {
		if file == fileToCompare {
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
			TraverseFolders(nextFolder, func(x *File) {
				if isEqualFile(fileToCompare, x) {
					contains = true
					file = x
					return
				}
			})
		}
	}

	return contains, file
}

func isEqualFile(a, b *File) bool {
	return a.Name == b.Name &&
		a.Size == b.Size &&
		a.LastModified == b.LastModified
}
