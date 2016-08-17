package main

import "fmt"

func main() {
	LearningString()
	LearningBoolean()
	Variable()
	Variable2()
	InputProgram()
}

// InputProgram demonstrate input value using fmt.Scanf
func InputProgram() {
	fmt.Print("Enter a number :")

	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}

// Const demonstrate how to create final variable in go-lang
func Const() {
	const x = "Eko Kurniawan"
	// x = "Khannedy" // you can not do that
}

// Variable2 demonstrate how to create variable in go-lang
func Variable2() {
	x := "Eko Kurniawan"
	fmt.Println(x)

	var y = "Khannedy"
	fmt.Println(y)
}

// Variable demonstrate how to create variable in go-lang
func Variable() {
	var x string

	x = "first"
	fmt.Println(x)

	x += "second"
	fmt.Println(x)

	x = "first"
	var y = "second"

	fmt.Println(x == y)

	x = "hello"
	y = "hello"
	fmt.Println(x == y)
}

// LearningBoolean demonstrate how to use bool in go-lang
func LearningBoolean() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

// LearningString demonstrate how to using string in go-lang
func LearningString() {
	fmt.Println("Hello " + "World")
	fmt.Println("Hello World"[1])
	fmt.Println(len("Hello World"))
}
