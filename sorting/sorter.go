package sorting

import "git.rickiekarp.net/rickie/tree2yaml/model"

// ByName implements sort.Interface based on the Name field.
type ByName []*model.File

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// ByFolderName implements sort.Interface based on the Name field.
type ByFolderName []*model.Folder

func (a ByFolderName) Len() int           { return len(a) }
func (a ByFolderName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByFolderName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
