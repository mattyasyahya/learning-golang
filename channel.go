package main

import (
	"fmt"
	"time"
)

func ponger(c chan<- string) {
	for i := 0; i < 10; i++ {
		fmt.Println("send pong")
		c <- "pong"
	}
}

func pinger(c chan<- string) {
	for i := 0; i < 10; i++ {
		fmt.Println("send ping")
		c <- "ping"
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println("receive", msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
