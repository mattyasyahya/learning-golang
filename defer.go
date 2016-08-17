package main

import "fmt"

func first(){
	fmt.Println("First")
}

func second(){
	fmt.Println("Second")
}

func third(){
	defer second()
	first()
}

func main() {
	third()
}
