package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	result, _ := bcrypt.GenerateFromPassword([]byte("rahasia"), 10)

	fmt.Println(string(result))

	error := bcrypt.CompareHashAndPassword(result, []byte("rahasia"))
	fmt.Println(error)
}
