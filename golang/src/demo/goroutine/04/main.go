package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 向管道中写入1-120000个数
func putNum(intChan chan int) {
	for i := 2; i < 120000; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

// 从管道中取数据，并判断是否是素数并放入到另外的管道中
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	// close(primeChan) // 如果一个channel关闭了就不能向这个channel中再写入数据了
	exitChan <- true // 协程完毕后往exitChan管道中写入一个标识
	wg.Done()
}

func printPrime(primeChan chan int) {
	// for v := range primeChan {
	// 	fmt.Println(v)
	// }
	wg.Done()
}

func main() {
	start := time.Now().UnixNano() / 1e6 // 毫秒
	// 使用goroutine channel统计1-120000之间的素数
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 50000) // 储存素数的管道，如果计算的数值比较多，primeChan的容量要大一些，否则会出现死锁
	exitChan := make(chan bool, 16)    // 协程数量

	wg.Add(1)
	go putNum(intChan)

	// 开启16个协程进行素数计算
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}
	wg.Add(1)
	go printPrime(primeChan)

	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		// 当exitChan管道中的标识全部被取出则关闭primeChan管道
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()
	end := time.Now().UnixNano() / 1e6 // 毫秒
	fmt.Println("素数的数量:", len(primeChan))
	fmt.Println("结束....", end-start, "毫秒")
}
