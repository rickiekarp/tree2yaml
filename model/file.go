package model

import (
	"fmt"

	"git.rickiekarp.net/rickie/tree2yaml/hash"
)

func (k File) Hash() uint64 {
	return hash.CalcFnv([]byte(fmt.Sprintf("%s-%d-%s-%d-%d-%s",
		k.Name, k.Size, k.LastModified.String(), k.Crc32, k.Crc64, k.Md5)))
}
