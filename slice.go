package main

import "fmt"

func main() {
	Append()
	Copy()
	Slice()
}

// Slice demonstrate how to using slice
func Slice() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(data[2:7])
	fmt.Println(data[:7])
	fmt.Println(data[2:])
	fmt.Println(data[:])
}

// Copy demonstrate how to copy data to int
func Copy() {
	first := []int{1, 2, 3, 4}
	second := make([]int, 10)
	copy(second, first)
	fmt.Println(second)
}

// Append demonstrate how to append data to slice
func Append() {
	slice := []int{1, 2, 3, 4}
	slice = append(slice, 1, 2, 3, 4)
	fmt.Println(slice)
}
