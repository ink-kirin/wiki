package main

import (
	"fmt"
	"reflect"
)

func fn(x interface{}) {
	// reflect.Int64 reflect.Float64 reflect.String reflect.Bool
	v := reflect.ValueOf(x) // 通过反射获取变量的值
	fmt.Printf("%T \n", v)  // reflect.Value
	fmt.Println(v.Kind())   // int
	vm := v.Int() + 2       // 获取反射对应类型的原始值
	fmt.Println(vm)         // 14
}

func main() {
	fn(12)
}
