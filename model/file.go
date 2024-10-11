package model

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func (k File) Sha1() string {
	s := fmt.Sprintf("%d-%d", k.Size, k.LastModified.Unix())
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}
