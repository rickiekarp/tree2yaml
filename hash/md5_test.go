package hash

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCalcMd5(t *testing.T) {
	// Create a temporary file and write some content to it
	content := []byte("Hello, world!")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Calculate the MD5 hash of the file
	expected := "6cd3556deb0da54bca060b4c39479839"
	actual := CalcMd5(tmpfile.Name())

	// Compare the actual result with the expected result
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}
