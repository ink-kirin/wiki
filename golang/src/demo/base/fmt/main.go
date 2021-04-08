package main

import "fmt"

func main() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	/*通常Printf格式化字符串包含多个%参数时将会包含对应相同数量的额外操作数，
	但是%之后的 [1] 副词告诉Printf函数再次使用第一个操作数。
	第二，%后的#副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀。*/

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"

	var (
		name = "simba"
		age  = 30
	)
	fmt.Printf("name=%s age=%d \n", name, age) // %s字符串输出  %d十进制输出
	fmt.Printf("name=%T age=%T \n", name, age) // %T输出变量类型

	var num1 = 99
	fmt.Printf("num1=%v num1=%T \n", num1, num1) // %v原样输出

	var num2 = 2.3333222
	fmt.Printf("num2=%v 类型:%[1]T \n", num2)

}
