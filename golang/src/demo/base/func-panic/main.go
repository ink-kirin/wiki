package main

import (
	"errors"
	"fmt"
)

/*
Go 1.13 之前没有异常机制的，使用 panic/recover模式处理错误， 1.13之后使用errors
panic可以在任意地方使用，但recover只有在defer调用的函数中有效
*/
func fn() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	panic("抛出一个异常") // panic 抛出一个异常不继续执行
}

func fn1(x, y int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err) // err: runtime error: integer divide by zero
		}
	}()
	return x / y
}

func readFile(file string) error {
	if file == "main.go" {
		return nil
	}
	return errors.New("文件读取失败")
}

func myFn() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("发送邮件通知", err)
		}
	}()
	err := readFile("xxx")
	if err != nil {
		panic(err)
	}
}

func main() {
	fn()
	fmt.Println("success")

	fmt.Println("fn1:", fn1(1, 0)) // 0

	myFn()
	fmt.Println("继续执行")
}
