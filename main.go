package main

import "fmt"

func main() {
  LearningString()
  LearningBoolean()
  Variable()
  Variable2()
  InputProgram()
}

func InputProgram()  {
  fmt.Print("Enter a number :")

  var input float64
  fmt.Scanf("%f", &input)

  output := input * 2
  fmt.Println(output)
}

func Const()  {
  const x = "Eko Kurniawan"
  // x = "Khannedy" // you can not do that
}

func Variable2()  {
  x := "Eko Kurniawan"
  fmt.Println(x)

  var y = "Khannedy"
  fmt.Println(y)
}

func Variable()  {
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

func LearningBoolean()  {
  fmt.Println(true && true)
  fmt.Println(true && false)
  fmt.Println(true || true)
  fmt.Println(true || false)
  fmt.Println(!true)
}

func LearningString(){
  fmt.Println("Hello " + "World")
  fmt.Println("Hello World"[1])
  fmt.Println(len("Hello World"))
}
