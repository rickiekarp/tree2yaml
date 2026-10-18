package generator

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	"git.rickiekarp.net/rickie/tree2yaml/model"
	"gopkg.in/yaml.v2"
)

func TestBuildTree(t *testing.T) {
	root := t.TempDir()

	if err := os.WriteFile(filepath.Join(root, "root.txt"), []byte("12345"), 0644); err != nil {
		t.Fatalf("write root file: %v", err)
	}
	sub := filepath.Join(root, "sub")
	if err := os.Mkdir(sub, 0755); err != nil {
		t.Fatalf("mkdir sub: %v", err)
	}
	if err := os.WriteFile(filepath.Join(sub, "sub.txt"), []byte("abc"), 0644); err != nil {
		t.Fatalf("write sub file: %v", err)
	}

	tree := buildTree(root)
	if tree == nil {
		t.Fatal("buildTree returned nil")
	}

	if tree.RootDir != path.Clean(root) {
		t.Errorf("RootDir = %q, want %q", tree.RootDir, path.Clean(root))
	}
	if tree.Size != 8 {
		t.Errorf("Size = %d, want 8", tree.Size)
	}
	if tree.ParserVersion != Version {
		t.Errorf("ParserVersion = %q, want %q", tree.ParserVersion, Version)
	}
	if tree.Tree == nil {
		t.Fatal("Tree is nil")
	}
	if tree.Tree.Name != path.Base(root) {
		t.Errorf("Tree.Name = %q, want %q", tree.Tree.Name, path.Base(root))
	}
	if len(tree.Tree.Files) != 1 || tree.Tree.Files[0].Name != "root.txt" {
		t.Errorf("Tree.Files = %+v, want [root.txt]", tree.Tree.Files)
	}
	if len(tree.Tree.Folders) != 1 || tree.Tree.Folders[0].Name != "sub" {
		t.Fatalf("Tree.Folders = %+v, want [sub]", tree.Tree.Folders)
	}
	if len(tree.Tree.Folders[0].Files) != 1 || tree.Tree.Folders[0].Files[0].Name != "sub.txt" {
		t.Errorf("sub.Files = %+v, want [sub.txt]", tree.Tree.Folders[0].Files)
	}
}

func TestMarshalFileTree(t *testing.T) {
	tree := &model.FileTree{
		RootDir:       "/root",
		ParserVersion: "test",
		Size:          3,
		Tree: &model.Folder{
			Name:  "root",
			Files: []*model.File{{Name: "a", Size: 3}},
		},
	}

	data, err := marshalFileTree(tree)
	if err != nil {
		t.Fatalf("marshalFileTree: %v", err)
	}

	var got model.FileTree
	if err := yaml.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal output: %v", err)
	}
	if got.RootDir != "/root" || got.Size != 3 {
		t.Errorf("round-trip = %+v", got)
	}
	if got.Tree == nil || len(got.Tree.Files) != 1 || got.Tree.Files[0].Name != "a" {
		t.Errorf("round-trip Tree = %+v", got.Tree)
	}
}

func TestWriteFiletreeToFile(t *testing.T) {
	tree := &model.FileTree{RootDir: "/root", Tree: &model.Folder{Name: "root"}}
	out := filepath.Join(t.TempDir(), "out.yaml")

	writeFiletreeToFile(tree, out)

	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	if len(data) == 0 {
		t.Error("output file is empty")
	}
}
