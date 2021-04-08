package main

import (
	"fmt"
	"reflect"
)

type myInt int

// Info 结构体
type Info struct {
	Name string
	Age  int
}

// 空接口可以接受任意类型的参数
func fn(x interface{}) {
	v := reflect.TypeOf(x) // 获取变量的类型
	fmt.Printf("变量类型:%v 名称:%v 种类:%v \n", v, v.Name(), v.Kind())
}

func main() {
	a := 12
	b := 23.11
	c := "golang"
	d := false
	fn(a)
	fn(b)
	fn(c)
	fn(d)

	var e myInt = 21
	fn(e)

	var f = Info{
		Name: "simba",
		Age:  32,
	}
	fn(f)

	var g = 2
	fn(&g) // 指针类型 变量类型:*int 名称: 种类:ptr

	var h = [3]int{1, 2, 3} // array
	fn(h)                   // 变量类型:[3]int 名称: 种类:array

	var i = []string{"simba", "kiri"} // slice
	fn(i)                             // 变量类型:[]string 名称: 种类:slice

}