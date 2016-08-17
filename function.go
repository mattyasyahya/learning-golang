package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, value := range xs {
		total += value
	}
	return total / float64(len(xs))
}

func variadicParameter(args ...int) int {
	total := 0
	for _, value := range args {
		total += value
	}
	return total
}

func returnVariable() (r int) {
	r = 10
	return
}

func multipleValue() (int, int) {
	return 1, 2
}

func closure() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 2))
}

func incrementNumber() {
	x := 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}

func main() {
	result := average([]float64{1, 2, 3, 4, 4, 5})
	fmt.Println(result)
	fmt.Println(multipleValue())
	fmt.Println(returnVariable())
	fmt.Println(variadicParameter(1, 2, 3, 4, 5, 6, 7, 8, 9, 0))

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(variadicParameter(slice...))

	closure()
	incrementNumber()
}
