package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Eko Kurniawan Khannedy"
	fmt.Println(strings.Contains(strings.ToLower(name), "eko"))

	fmt.Println(len(name))
	fmt.Println(strings.Count(name, "a"))
}
