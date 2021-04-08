package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 只写管道
func fn1(ch chan<- int) {
	for i := 0; i < cap(ch); i++ {
		ch <- i
		fmt.Printf("【写入】%v成功\n", i)
		time.Sleep(time.Millisecond * 50)
	}
	close(ch)
	wg.Done()
}

// 只读管道
func fn2(ch <-chan int) {
	for v := range ch {
		fmt.Printf("【读取】%v成功\n", v)
	}
	wg.Done()
}

func main() {
	// 写入和读取并行
	ch := make(chan int, 10)
	wg.Add(1)
	go fn1(ch)
	wg.Add(1)
	go fn2(ch)
	wg.Wait()
	fmt.Printf("结束...")
}
