package main

import "fmt"

// 交换函数
func swap(a, b *int) {
	// 取a指针的值, 赋给临时变量t
	t := *a
	// 取b指针的值, 赋给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = t
}

// 指针
func main() {
	// 指针地址和指针类型
	// Go语言中使用在变量名前面添加&操作符（前缀）来获取变量的内存地址（取地址操作）
	// 其中 v 代表被取地址的变量，变量 v 的地址使用变量 ptr 进行接收，ptr 的类型为*T，称做 T 的指针类型，*代表指针。
	var v int = 1
	ptr := &v // v 的类型为 T
	// 使用 fmt.Printf 的动词%p打印变量的内存地址
	fmt.Printf("%p \n", ptr)

	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, ptr类型为*string
	potr := &house
	// 打印ptr的类型
	fmt.Printf("potr type: %T\n", potr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", potr)
	// 对指针进行取值操作
	value := *potr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

	// 修改指针的值
	// 准备两个变量, 赋值1和2
	x, y := 1, 2
	// 交换变量值
	swap(&x, &y)
	// 输出变量值
	fmt.Println(x, y)

	// 使用new创建指针 new(类型)
	str := new(string) // 指针类型
	*str = "Go语言教程"
	fmt.Println(*str)

	// 创建指针 （推荐使用）
	var sv int = 1
	pt := &sv
	fmt.Println(*pt)
}
