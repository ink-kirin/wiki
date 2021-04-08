package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var wg sync.WaitGroup
var mutex sync.Mutex
var rwmutex sync.RWMutex

func fn() {
	defer wg.Done()
	mutex.Lock()
	count++
	mutex.Unlock()
	fmt.Println(count)
}

var m = make(map[int]int, 0)

func fn1(num int) {
	// 如果不加互斥锁会出现资源竞争的问题
	mutex.Lock()
	var sum = 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	m[num] = sum
	fmt.Println(num, sum)
	mutex.Unlock()
	wg.Done()
}

func write() {
	mutex.Lock()
	fmt.Println("写操作")
	time.Sleep(time.Second) // 停止1秒
	mutex.Unlock()
	wg.Done()
}

func read(n int) {
	rwmutex.RLock()
	fmt.Println("读操作--", n)
	rwmutex.RUnlock()
	wg.Done()
}

func main() {
	// 互斥锁 (并行变为串行)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go fn()
	}
	wg.Wait()
	fmt.Println("【fn】结束....")

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go fn1(i)
	}
	wg.Wait()
	fmt.Println("【fn1】结束....")

	// 读写互斥锁(读写锁)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go read(i)
	}
	wg.Wait()
	fmt.Println("【write read】结束....")
}
