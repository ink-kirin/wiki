package main

import "fmt"

type A interface{} // 空接口 表示没有任何约束 任意类型都可以实现空接口

// 用空接口做参数表示可以接受任意类型的变量
func show(i interface{}) {
	fmt.Printf("%v %[1]T \n", i)
}

// 空接口的使用
func main() {
	var a A
	var str = "golang"
	a = str
	fmt.Printf("%v %[1]T \n", a)

	var num = 2
	a = num
	fmt.Printf("%v %[1]T \n", a)

	// 空接口可以当做类型使用，可以表示任意类型
	var i interface{}
	i = 20
	fmt.Printf("%v %[1]T \n", i)
	i = "haha"
	fmt.Printf("%v %[1]T \n", i)

	var m = make(map[string]interface{}) // map 值为任意类型
	m["name"] = "simba"
	m["age"] = 20
	m["height"] = 1.72
	fmt.Printf("%v %[1]T \n", m) // map[age:20 height:1.72 name:simba] map[string]interface {}

	var s = []interface{}{1, 2, 3.44, "kirin", true} // 任意类型值的切片
	fmt.Printf("%v %[1]T \n", s)                     // [1 2 3.44 kirin true] []interface {}

}
