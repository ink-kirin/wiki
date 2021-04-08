package main

import (
	"fmt"
	"time"
)

func speed1(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "speed 1"
}

func speed2(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "speed 2"
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go speed1(c1)
	go speed2(c2)

	fmt.Println("The first to arrive is:")

	select {
	case s1 := <-c1:
		fmt.Println(s1)
	case s2 := <-c2:
		fmt.Println(s2)
	}
}
