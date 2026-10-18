package extractor

import (
	"testing"
	"time"

	"git.rickiekarp.net/rickie/tree2yaml/model"
)

func buildTree() (*model.Folder, *model.Folder, *model.Folder) {
	root := &model.Folder{Name: "root"}
	a := root.AddFolder("a")
	b := a.AddFolder("b")
	c := b.AddFolder("c")
	return root, b, c
}

func TestFindFolderSinglePathReturnsRoot(t *testing.T) {
	root, _, _ := buildTree()
	if got := FindFolder(root, []string{"only"}); got != root {
		t.Errorf("FindFolder = %+v, want root", got)
	}
}

func TestFindFolderNested(t *testing.T) {
	root, b, c := buildTree()

	if got := FindFolder(root, []string{"a", "b"}); got != b {
		t.Errorf("FindFolder([a b]) = %+v, want b", got)
	}
	if got := FindFolder(root, []string{"a", "b", "c"}); got != c {
		t.Errorf("FindFolder([a b c]) = %+v, want c", got)
	}
}

func TestFindFolderMissing(t *testing.T) {
	root, _, _ := buildTree()
	if got := FindFolder(root, []string{"nope", "missing"}); got != nil {
		t.Errorf("FindFolder = %+v, want nil", got)
	}
}

func TestFindSubDirectory(t *testing.T) {
	_, b, _ := buildTree()
	if got := findSubDirectory(b.Folders, "c"); got == nil || got.Name != "c" {
		t.Errorf("findSubDirectory = %+v, want c", got)
	}
	if got := findSubDirectory(b.Folders, "missing"); got != nil {
		t.Errorf("findSubDirectory = %+v, want nil", got)
	}
}

func TestFilterByModDate(t *testing.T) {
	cutoff := time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)
	newer := &model.File{Name: "new", LastModified: cutoff.Add(24 * time.Hour)}
	older := &model.File{Name: "old", LastModified: cutoff.Add(-24 * time.Hour)}

	newFolder := func() *model.Folder {
		return &model.Folder{Name: "root", Files: []*model.File{newer, older}}
	}

	t.Run("new direction", func(t *testing.T) {
		got := FilterByModDate(newFolder(), "2023-06-15", "new")
		if len(got.Files) != 1 || got.Files[0].Name != "new" {
			t.Errorf("Files = %+v, want [new]", got.Files)
		}
	})

	t.Run("old direction", func(t *testing.T) {
		got := FilterByModDate(newFolder(), "2023-06-15", "old")
		if len(got.Files) != 1 || got.Files[0].Name != "old" {
			t.Errorf("Files = %+v, want [old]", got.Files)
		}
	})

	t.Run("default direction filters new", func(t *testing.T) {
		got := FilterByModDate(newFolder(), "2023-06-15", "")
		if len(got.Files) != 1 || got.Files[0].Name != "new" {
			t.Errorf("Files = %+v, want [new]", got.Files)
		}
	})

	t.Run("invalid date returns folder unchanged", func(t *testing.T) {
		got := FilterByModDate(newFolder(), "not-a-date", "new")
		if len(got.Files) != 2 {
			t.Errorf("Files = %+v, want unchanged 2 files", got.Files)
		}
	})
}
