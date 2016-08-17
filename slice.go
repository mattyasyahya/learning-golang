package main

import "fmt"

func main() {
	Append()
	Copy()
}

func Copy() {
	first := []int{1, 2, 3, 4}
	second := make([]int, 10)

	copy(second, first)
	fmt.Println(second)
}

func Append() {
	slice := []int{1, 2, 3, 4}

	slice = append(slice, 1, 2, 3, 4)

	fmt.Println(slice)
}
