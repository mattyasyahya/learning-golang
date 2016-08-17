package main

import "fmt"

func main() {
	samplePanic()
}

func recoverPanic() {
	str := recover()
	fmt.Println(str)
}

func samplePanic() {
	defer recoverPanic()
	array := []int{1, 2, 3, 4}
	fmt.Println(array[10])
}
