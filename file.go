package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	complexWay()
	easyWay()
	createFile()
	walk()
}

func walk() {
	filepath.Walk(".", func(path string, info os.FileInfo, error error) error {
		fmt.Println(path)
		return nil
	})
}

func createFile() {
	file, error := os.Create("output/text.txt")
	if error != nil {
		panic(error)
	}
	defer file.Close()

	file.WriteString("Hello World")
}

func easyWay() {
	bs, error := ioutil.ReadFile("input/text.txt")
	if error != nil {
		panic(error)
	}

	str := string(bs)
	fmt.Println(str)
}

func complexWay() {
	file, error := os.Open("input/text.txt")
	if error != nil {
		panic(error)
	}
	defer file.Close()

	stat, error := file.Stat()
	if error != nil {
		panic(error)
	}

	value := make([]byte, stat.Size())
	_, error = file.Read(value)
	if error != nil {
		panic(error)
	}

	str := string(value)
	fmt.Println(str)
}
