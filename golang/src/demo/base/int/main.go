package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 = 128 // 默认int类型 64位系统下int64 32位系统下int32
	fmt.Println(num1)

	var num2 int8 = 89
	// num2 = 339 // 超出int8对范围
	fmt.Printf("num2=%v 类型:%T \n", num2, num2)
	fmt.Println(unsafe.Sizeof(num2)) // 内存占用空间（字节）

	var num3 int32 = 98
	var num4 int64 = 99
	fmt.Println(num3 + int32(num4)) // 类型转换 同类型才能做计算等操作

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	fmt.Printf("KB=%v MB=%v GB=%v TB=%v PB=%v \n", KB, MB, GB, TB, PB)

	// 位移运算
	// a << b (a左移b位后的值): a*2^b
	// 9 << 4 9*2^4 = 144

	f := 3.141 // a float64
	i := int(f)
	fmt.Println(f, i) // "3.141 3"
	f = 1.99
	fmt.Println(int(f)) // "1"

	f = 1e100
	i = int(f)
	fmt.Println(i)

	// 交换两个变量的值
	a := 20
	b := 34
	// 方法一
	var c = a
	a = b
	b = c
	fmt.Println(a, b)
	// 方法二
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)

}
