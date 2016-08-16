package main

import "fmt"

func main() {
  LearningString()
  LearningBoolean()
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
