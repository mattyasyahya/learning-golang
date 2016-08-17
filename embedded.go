package main

import "fmt"

func main() {
	fmt.Println("Start program")

	android := new(Android)
	android.Name = "Eko"
	android.talk()
}

// Person type
type Person struct {
	Name string
}

func (p *Person) talk() {
	fmt.Println("I'm " + p.Name + " talking")
}

// Android type
type Android struct {
	Person
	Model string
}
