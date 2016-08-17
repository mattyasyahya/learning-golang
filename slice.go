package main

import "fmt"

func main() {
	Append()
	Copy()
	Slice()
}

func Slice() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(data[2:7])
	fmt.Println(data[:7])
	fmt.Println(data[2:])
	fmt.Println(data[:])
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
