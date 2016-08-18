package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func crc() {
	h := crc32.NewIEEE()
	h.Write([]byte("Eko Kurniawan"))
	r := h.Sum32()
	fmt.Println(r)
}

func getHash(filename string) (uint32, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	h := crc32.NewIEEE()

	_, err = io.Copy(h, f)
	if err != nil {
		return 0, err
	}

	return h.Sum32(), nil
}

func main() {
	crc()

	h1, err := getHash("input/text.txt")
	if err != nil {
		return
	}

	h2, err := getHash("output/text.txt")
	if err != nil {
		return
	}

	fmt.Println(h1, h2, h1 == h2)
}
