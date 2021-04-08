package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 30)
	for i := 0; i < 30; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "simba-" + fmt.Sprintf("%d", i) // int 转换 string 进行拼接
	}

	for {
		select {
		case v := <-intChan:
			fmt.Println("intChan:", v)
			time.Sleep(time.Millisecond * 50) // 停止50毫秒
		case v := <-stringChan:
			fmt.Println("stringChan:", v)
			time.Sleep(time.Millisecond * 50) // 停止50毫秒
		default:
			fmt.Println("结束....")
			return
		}
	}
}
