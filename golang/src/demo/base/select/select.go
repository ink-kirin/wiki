package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	w := make(chan int)
	go worker(id, w)
	return w
}

func main() {
	var c1, c2 = generator(), generator()
	worker := createWorker(0)
	var values []int
	tm := time.After(10 * time.Second) // 10秒后返回chan
	tick := time.Tick(time.Second)     // 每一秒返回一个chan
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond): // 800ms返回一个chan
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("values len:", len(values))
		case <-tm:
			fmt.Println("bb")
			return
		}
	}

}
