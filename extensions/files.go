package extensions

import (
	"os"
)

func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	} else {
		return false
	}
}
