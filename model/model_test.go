package model

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestIsFile(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "a.txt")
	if err := os.WriteFile(filePath, []byte("hello"), 0644); err != nil {
		t.Fatalf("write file: %v", err)
	}

	if ok, err := IsFile(filePath); err != nil || !ok {
		t.Errorf("IsFile(file) = %v, %v; want true, nil", ok, err)
	}
	if ok, err := IsFile(dir); err != nil || ok {
		t.Errorf("IsFile(dir) = %v, %v; want false, nil", ok, err)
	}
	if _, err := IsFile(filepath.Join(dir, "missing")); err == nil {
		t.Error("IsFile(missing) expected error, got nil")
	}
}

func TestFolderAddFile(t *testing.T) {
	root := &Folder{Name: "root"}
	a := &File{Name: "a"}
	b := &File{Name: "b"}

	root.AddFile(a)
	root.AddFile(b)

	if len(root.Files) != 2 {
		t.Fatalf("len(Files) = %d, want 2", len(root.Files))
	}
	if root.Files[0] != a || root.Files[1] != b {
		t.Errorf("Files = %+v, want [a b]", root.Files)
	}
}

func TestFolderAddFolder(t *testing.T) {
	root := &Folder{Name: "root"}
	sub := root.AddFolder("sub")

	if sub.Name != "sub" {
		t.Errorf("Name = %q, want %q", sub.Name, "sub")
	}
	if len(root.Folders) != 1 || root.Folders[0] != sub {
		t.Errorf("Folders = %+v, want [sub]", root.Folders)
	}
}

func TestTraverseFiles(t *testing.T) {
	files := []*File{{Name: "a"}, {Name: "b"}, {Name: "c"}}
	var visited []string
	TraverseFiles(files, func(f *File) { visited = append(visited, f.Name) })

	if len(visited) != 3 || visited[0] != "a" || visited[2] != "c" {
		t.Errorf("visited = %v, want [a b c]", visited)
	}
}

func TestTraverseFilesAndFolders(t *testing.T) {
	root := &Folder{Name: "root"}
	root.AddFile(&File{Name: "rootfile"})
	sub := root.AddFolder("sub")
	sub.AddFile(&File{Name: "subfile"})
	sub.AddFolder("deeper")

	var files, folders []string
	TraverseFilesAndFolders(
		root,
		func(f *File) { files = append(files, f.Name) },
		func(f *Folder) { folders = append(folders, f.Name) },
	)

	if len(files) != 2 || files[0] != "rootfile" || files[1] != "subfile" {
		t.Errorf("files = %v, want [rootfile subfile]", files)
	}
	if len(folders) != 2 || folders[0] != "sub" || folders[1] != "deeper" {
		t.Errorf("folders = %v, want [sub deeper]", folders)
	}
}

func TestTraverseNilCallbacks(t *testing.T) {
	root := &Folder{Name: "root", Files: []*File{{Name: "a"}}, Folders: []*Folder{{Name: "sub"}}}
	TraverseFiles(root.Files, nil)
	TraverseFolder(root.Folders, nil)
	TraverseFilesAndFolders(root, nil, nil)
}

func TestLoadFilelist(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "list.yaml")
	data := `directory: /root
parserVersion: test
size: 10
tree:
  name: root
  files:
    - name: a.txt
      size: 5
      lastmodified: 2020-01-02T03:04:05Z
      crc32: 123
`
	if err := os.WriteFile(path, []byte(data), 0644); err != nil {
		t.Fatalf("write yaml: %v", err)
	}

	tree := LoadFilelist(path)
	if tree == nil {
		t.Fatal("LoadFilelist returned nil")
	}
	if tree.RootDir != "/root" {
		t.Errorf("RootDir = %q, want %q", tree.RootDir, "/root")
	}
	if tree.ParserVersion != "test" {
		t.Errorf("ParserVersion = %q, want %q", tree.ParserVersion, "test")
	}
	if tree.Size != 10 {
		t.Errorf("Size = %d, want 10", tree.Size)
	}
	if tree.Tree == nil || len(tree.Tree.Files) != 1 {
		t.Fatalf("Tree.Files = %+v, want 1 file", tree.Tree)
	}
	if tree.Tree.Files[0].Name != "a.txt" {
		t.Errorf("file name = %q, want %q", tree.Tree.Files[0].Name, "a.txt")
	}
	want := time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)
	if !tree.Tree.Files[0].LastModified.Equal(want) {
		t.Errorf("LastModified = %v, want %v", tree.Tree.Files[0].LastModified, want)
	}
}
