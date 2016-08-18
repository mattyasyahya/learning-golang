package main

import (
	"container/list"
	"fmt"
	"sort"
)

func sampleList() {
	numbers := list.New()
	numbers.PushBack(1)
	numbers.PushBack("eko")
	numbers.PushBack(2)

	for value := numbers.Front(); value != nil; value = value.Next() {
		fmt.Println(value.Value)
	}
}

// Person type
type Person struct {
	Name string
	Age  int
}

// ByName for slice of Person
type ByName []Person

func (data ByName) Len() int {
	return len(data)
}

func (data ByName) Less(i, j int) bool {
	return data[i].Name < data[j].Name
}

func (data ByName) Swap(i, j int) {
	data[i], data[j] = data[j], data[i]
}

// ByAge for slice Person
type ByAge []Person

func (data ByAge) Len() int {
	return len(data)
}

func (data ByAge) Less(i, j int) bool {
	return data[i].Age < data[j].Age
}

func (data ByAge) Swap(i, j int) {
	data[i], data[j] = data[j], data[i]
}

func sampleSort() {
	persons := []Person{
		{"Kurniawan", 29},
		{"Khannedy", 30},
		{"Eko", 28},
	}

	sort.Sort(ByName(persons))
	fmt.Println(persons)

	sort.Sort(ByAge(persons))
	fmt.Println(persons)
}

func main() {
	sampleList()
	sampleSort()
}
