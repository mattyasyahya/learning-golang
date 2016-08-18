package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("Start program")
	var data bytes.Buffer
	data.Write([]byte("Eko Kurniawan Khannedy"))

	fmt.Println(string(data.Bytes()))
}
