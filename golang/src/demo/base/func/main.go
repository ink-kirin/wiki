package main

import "fmt"

type cal func(int, int) int // 表示定义一个cal的方法类型

func add(x, y int) int {
	return x + y
}

func jisuan(x, y int, op cal) int {
	return op(x, y)
}

// 函数作为返回值
func do(o string) cal {
	switch o {
	case "+":
		return add
	case "-":
		return func(x, y int) int {
			return x - y
		}
	default:
		return nil
	}
}

// 递归函数
func fn(n int) int {
	if n > 1 {
		return n + fn(n-1)
	}
	return 1
}

// 可变参数
func sumFn(x ...int) int {
	// x 切片
	// fmt.Println(x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// 返回多个值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 定义返回值命名
func calc1(x, y int) (sum int, sub int) {
	// fmt.Println(sum, sub) // 0 0 初始值
	sum = x + y
	sub = x - y
	return
}

func main() {
	sum := sumFn(1, 2, 3, 4)
	fmt.Println(sum) // 10

	a, b := calc(1, 3)
	fmt.Println(a, b) // 4 -2

	a1, b1 := calc1(3, 2)
	fmt.Println(a1, b1) // 5 1

	// 使用一个函数作为另外一个函数的参数
	sum1 := jisuan(3, 5, add)
	fmt.Println(sum1) // 8

	// 匿名函数
	sum2 := jisuan(2, 5, func(x, y int) int {
		return x * y
	})
	fmt.Println(sum2) // 10

	sum3 := do("+")
	fmt.Println(sum3(3, 9)) // 12

	// 匿名函数接受参数 直接执行
	func(x, y int) {
		fmt.Println(x * y) // 24
	}(8, 3)
}
