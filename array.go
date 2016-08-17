package main

import "fmt"
import "math/rand"

func main() {
	Array()
	ShortArray()
}

// ShortArray demonstrate how to user array in go-lang
func ShortArray() {
	array := [5]int{
		1, 2, 3, 4, 5,
	}

	for _, value := range array {
		fmt.Println(value)
	}
}

// Array demonstrate how to use array
func Array() {
	var array [10]int

	array[0] = rand.Intn(10)
	array[1] = rand.Intn(10)
	array[2] = rand.Intn(10)
	array[3] = rand.Intn(10)
	array[4] = rand.Intn(10)
	array[5] = rand.Intn(10)
	array[6] = rand.Intn(10)
	array[7] = rand.Intn(10)
	array[8] = rand.Intn(10)
	array[9] = rand.Intn(10)

	fmt.Println(array)

	var total int

	for i := 0; i < len(array); i++ {
		total += array[i]
	}

	fmt.Println(total)

	total = 0
	for _, value := range array {
		total += value
	}

	fmt.Println(total)

}
