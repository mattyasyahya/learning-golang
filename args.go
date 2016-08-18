package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	maxp := flag.Int("max", 5, "the max value")

	var seed = int64(*maxp)

	rand.Seed(seed)
	flag.Parse()

	fmt.Println(rand.Intn(*maxp))
	fmt.Println(rand.Intn(*maxp))
	fmt.Println(rand.Intn(*maxp))
	fmt.Println(rand.Intn(*maxp))
}
