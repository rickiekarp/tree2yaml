package model

func (f *FileTree) ContainsFile(fileToCompare *File) bool {
	var contains = false

	for _, file := range f.Tree.Files {
		if file == fileToCompare {
			return true
		}
	}

	for _, folder := range f.Tree.Folders {
		TraverseFiles(folder.Files, func(x *File) {
			if isEqualFile(fileToCompare, x) {
				contains = true
				return
			}
		})

		for _, nextFolder := range folder.Folders {
			TraverseFolders(nextFolder, func(x *File) {
				if isEqualFile(fileToCompare, x) {
					contains = true
					return
				}
			})
		}
	}

	return contains
}

func isEqualFile(a, b *File) bool {
	return a.Name == b.Name &&
		a.Size == b.Size &&
		a.LastModified == b.LastModified
}
