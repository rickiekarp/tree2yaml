package hash

import (
	"bufio"
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"os"
)

func CalcCrc32(filePath string) uint32 {
	file, err := os.Open(filePath)
	if err != nil {
		os.Exit(1)
	}

	reader := bufio.NewReader(file)
	byteSlice, err := reader.Peek(reader.Size())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	table := crc32.MakeTable(crc32.IEEE)
	checksum := crc32.Checksum(byteSlice, table)
	return checksum
}

func CalcCrc64(filePath string) uint64 {
	file, err := os.Open(filePath)
	if err != nil {
		os.Exit(1)
	}

	reader := bufio.NewReader(file)
	byteSlice, err := reader.Peek(reader.Size())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	table := crc64.MakeTable(crc32.IEEE)
	checksum := crc64.Checksum(byteSlice, table)
	return checksum
}
