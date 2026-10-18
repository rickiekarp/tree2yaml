package printer

import (
	"io"
	"os"
	"strings"
	"testing"

	"git.rickiekarp.net/rickie/tree2yaml/model"
)

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("pipe: %v", err)
	}
	os.Stdout = w

	fn()

	if err := w.Close(); err != nil {
		t.Fatalf("close pipe: %v", err)
	}
	os.Stdout = old

	out, _ := io.ReadAll(r)
	return string(out)
}

func TestPrintFilelist(t *testing.T) {
	tree := &model.FileTree{
		RootDir: "/root",
		Tree:    &model.Folder{Name: "root", Files: []*model.File{{Name: "a"}}},
	}

	out := captureStdout(t, func() { PrintFilelist(tree) })

	if !strings.Contains(out, "directory: /root") {
		t.Errorf("output missing directory: %q", out)
	}
	if !strings.Contains(out, "name: a") {
		t.Errorf("output missing file name: %q", out)
	}
}
