package main

import "fmt"

// 闭包

func adder() func() int {
	var i = 10
	return func() int {
		return i + 1
	}
}

func adder1() func(int) int {
	var i = 10
	return func(y int) int {
		i += y
		return i
	}
}
func main() {
	var f = adder()  // 表示支持方法
	fmt.Println(f()) // 11

	var fn = adder1()
	fmt.Println(fn(10)) // 20
	fmt.Println(fn(10)) // 30
	fmt.Println(fn(10)) // 40
}
