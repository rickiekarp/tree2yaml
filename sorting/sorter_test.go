package sorting

import (
	"sort"
	"testing"

	"git.rickiekarp.net/rickie/tree2yaml/model"
)

func TestByName(t *testing.T) {
	files := []*model.File{{Name: "charlie"}, {Name: "alpha"}, {Name: "bravo"}}
	sort.Sort(ByName(files))

	want := []string{"alpha", "bravo", "charlie"}
	for i, f := range files {
		if f.Name != want[i] {
			t.Errorf("files[%d].Name = %q, want %q", i, f.Name, want[i])
		}
	}
}

func TestByFolderName(t *testing.T) {
	folders := []*model.Folder{{Name: "zulu"}, {Name: "alpha"}, {Name: "mike"}}
	sort.Sort(ByFolderName(folders))

	want := []string{"alpha", "mike", "zulu"}
	for i, f := range folders {
		if f.Name != want[i] {
			t.Errorf("folders[%d].Name = %q, want %q", i, f.Name, want[i])
		}
	}
}

func TestSortInterface(t *testing.T) {
	var _ sort.Interface = ByName(nil)
	var _ sort.Interface = ByFolderName(nil)

	byName := ByName{{Name: "b"}, {Name: "a"}}
	if byName.Len() != 2 {
		t.Errorf("Len() = %d, want 2", byName.Len())
	}
	if !byName.Less(1, 0) {
		t.Error("Less(1,0) = false, want true")
	}
	byName.Swap(0, 1)
	if byName[0].Name != "a" {
		t.Errorf("after Swap, [0].Name = %q, want %q", byName[0].Name, "a")
	}
}
