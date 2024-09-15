package model

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"git.rickiekarp.net/rickie/tree2yaml/hash"
)

func (k File) Hash() uint64 {
	return hash.CalcFnv([]byte(fmt.Sprintf("%s-%d-%s-%d-%d-%s",
		k.Name, k.Size, k.LastModified.String(), k.Crc32, k.Crc64, k.Md5)))
}

func (k File) Sha1() string {
	s := fmt.Sprintf("%d-%d", k.Size, k.LastModified.Unix())
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}
