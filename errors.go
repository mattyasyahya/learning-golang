package main

import (
	"errors"
	"fmt"
)

func test(i int) (int, error) {
	if i >= 0 {
		return i, nil
	}
	return i, errors.New("number is < 0")
}

func main() {
	i, err := test(-1)
	if err != nil {
		fmt.Println("ups: " + err.Error())
	} else {
		fmt.Println(i)
	}
}
