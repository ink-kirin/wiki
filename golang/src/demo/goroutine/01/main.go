package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func test(num int) {
	defer wg.Done() // 计数器减一
	for i := 0; i < 10; i++ {
		fmt.Printf("协程(%v)的第(%v)条数据\n", num, i)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 设置计数器为1
		go test(i)
	}
	wg.Wait() // 阻塞等待计数器归零
	fmt.Println("结束")
}
