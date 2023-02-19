package model

func isEqualFile(a, b *File) bool {
	return a != nil && b != nil &&
		a.Name == b.Name &&
		a.Size == b.Size &&
		a.LastModified.UTC() == b.LastModified.UTC()
}
