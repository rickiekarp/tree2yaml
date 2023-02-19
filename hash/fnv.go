package hash

import (
	"hash/fnv"
)

func CalcFnv(data []byte) uint64 {
	h := fnv.New64a()
	h.Write(data)
	return h.Sum64()
}
