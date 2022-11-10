package main

import (
	"fmt"
	"time"
)

func ping(c chan string) { // reserved for Channel : chan
	for i := 0; ; i++ {
		c <- "ping" // used for sending and receiving message through the channel
	}
}

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong" // used for sending and receiving message through the channel
	}
}

func show(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go ping(c)
	go show(c)
	go pong(c)

	var firstHand string
	fmt.Scanln(&firstHand)
}
