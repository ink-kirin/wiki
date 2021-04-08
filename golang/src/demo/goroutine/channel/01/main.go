package main

import "fmt"

func main() {
	// 创建管道
	ch := make(chan int, 3)

	// 向管道内写入数据
	ch <- 10

	// 从管道内获取数据
	a := <-ch
	fmt.Println(a)

	ch <- 12

	// 管道的值，容量和长度
	fmt.Printf("值：%v, 容量：%v, 长度：%v\n", ch, cap(ch), len(ch))

	// 管道类型（引用类型）
	ch1 := ch
	ch1 <- 22
	<-ch
	d := <-ch
	fmt.Println(d)

	// 管道阻塞
	ch2 := make(chan int, 1)
	ch2 <- 20
	// ch2 <- 23 // 阻塞了

	ch3 := make(chan int, 1)
	ch3 <- 20
	<-ch3
	// <-ch3 // 阻塞了

	// 遍历管道
	// for range 遍历需要关闭管道
	ch4 := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch4 <- i
	}
	close(ch4) // 关闭管道
	// 管道类型只有值没有key
	for v := range ch4 {
		fmt.Println(v)
	}

	// for 遍历可以不用关闭
	ch5 := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch5 <- i
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch5)
	}
}
