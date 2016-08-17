package main

import "fmt"

func main() {
	Int()
	StringString()
	ShortMap()
}

// ShortMap demonstrate create map short way
func ShortMap() {
	elements := map[string]string{
		"first_name":  "Eko",
		"middle_name": "Kurniawan",
		"last_name":   "Khannedy",
	}

	fmt.Println(elements)
}

// StringString demonstrate map with key string and value string
func StringString() {
	data := make(map[string]string)

	data["first_name"] = "Eko Kurniawan"
	data["last_name"] = "Khannedy"
	data["middle_name"] = "Wrong"

	fmt.Println(data)
	fmt.Println(len(data))

	delete(data, "middle_name")
	fmt.Println(data)
	fmt.Println(len(data))

	if _, ok := data["first_name"]; ok {
		fmt.Println("OKE")
	}
}

// Int demonstrate map with key int and value int
func Int() {
	data := make(map[int]int)

	data[0] = 1
	data[1] = 1
	data[2] = 1

	fmt.Println(data)
	fmt.Println(len(data))
}
