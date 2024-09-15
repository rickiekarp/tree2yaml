package hash

import (
	"crypto/sha1"
	"encoding/hex"
)

func CalcSha1(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}
