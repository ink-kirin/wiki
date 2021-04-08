package main

import (
	"fmt"
)
// 匿名返回值
func f1() int {
	var a int // 0
	defer func () {
		a++
	}()
	return a
}
// 命名返回值
func f2() (a int) {
	defer func () {
		a++
	}()
	return a
}

func f3() (x int) {
	defer func (y int)  {
		y++
	}(x) // defer注册要延时执行的函数时该函数的所有参数都需要确定其值
	return x // 0
}

func main() {
	// defer 在函数执行完毕之后执行 多个defer从下往上的顺序执行
	defer func () {
		fmt.Println("闭包")
	}()
	fmt.Println("开始")
	defer fmt.Println("defer")
	fmt.Println("结束")

	// defer在命名返回值和匿名返回值 函数中表现不一样
	fmt.Println(f1()) // 0
	fmt.Println(f2()) // 1
	fmt.Println(f3()) // 0
}