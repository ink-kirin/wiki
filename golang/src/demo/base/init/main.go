package main

import "fmt"

func init() {
	fmt.Println("init")
}

/*
	init() 先于main函数执行
	最后引入的包优先执行init()
*/
func main() {
	fmt.Println("main")
}
