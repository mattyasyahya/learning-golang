package main

import "fmt"

func zero(x *int) {
	*x = 100
}

func usingAnd() {
	x := 10
	zero(&x)
	fmt.Println(x)
}

func usingNew() {
	x := new(int)
	*x = 1
	zero(x)
	fmt.Println(*x)
}

func square(x *int) {
	*x = *x * *x
}

func main() {
	usingAnd()
	usingNew()

	x := 10
	square(&x)
	fmt.Println(x)
}
