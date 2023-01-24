package hash

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func CalcMd5(filePath string) string {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}
