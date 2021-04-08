package main

import (
	"fmt"
)

func main() {
	/*
		complex 构建复数
		real 返回复数的实部
		imag 返回复数的虚部
	*/
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // (-5+10i)
	fmt.Println(real(x * y))         // -5
	fmt.Println(imag(x * y))         // 10

}
