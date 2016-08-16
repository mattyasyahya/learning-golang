package main

import "fmt"

func main() {
  LearningString()
}

func LearningString(){
  fmt.Println("Hello " + "World")
  fmt.Println("Hello World"[1])
  fmt.Println(len("Hello World"))
}
