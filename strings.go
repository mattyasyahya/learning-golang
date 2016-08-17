package main

import (
	"fmt"
	"strings"
)

func binary() {
	firstName := []byte("Eko Kurniawan")
	lastName := string([]byte("Khannedy"))

	fmt.Println(string(firstName))
	fmt.Println(lastName)
}

func main() {
	name := "Eko Kurniawan Khannedy"
	fmt.Println(strings.Contains(strings.ToLower(name), "eko"))

	fmt.Println(len(name))
	fmt.Println(strings.Count(name, "a"))

	binary()
}
