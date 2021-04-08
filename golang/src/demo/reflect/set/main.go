package main

import (
	"fmt"
	"reflect"
)

// 变量为引用类型
func fn(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v.Kind())        // ptr
	fmt.Println(v.Elem().Kind()) // int64
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(90)
	}
}

func main() {
	// 通过反射修改变量的值
	var a int64 = 32
	fn(&a)
	fmt.Println(a)
}
