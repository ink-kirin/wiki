package main

import "fmt"

// 自定义类型
type myInt int

// 类型别名
type myFload = float64

func main() {
	var a myInt = 2
	fmt.Printf("%v %T\n", a, a) // 2 main.myInt

	var b myFload = 22.9
	fmt.Printf("%v %T\n", b, b) // 2.9 float64
}