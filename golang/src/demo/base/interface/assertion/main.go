package main

import (
	"fmt"
)

// Print 定义一个方法，传入不同的类型实现不同的功能
func Print(x interface{}) {
	switch x.(type) { // x.(type) 获取x的类型, 只能结合switch使用
	case string:
		fmt.Println("string")
		break
	case int:
		fmt.Println("int")
		break
	case bool:
		fmt.Println("bool")
		break
	default:
		fmt.Println("传入错误...")
		break
	}
}

// 断言
func main() {
	// 断言
	var a1 interface{}
	a1 = "as"
	v, ok := a1.(string)
	if ok {
		fmt.Printf("%v \n", v)
	} else {
		fmt.Println("断言失败")
	}

	Print(true)

}
