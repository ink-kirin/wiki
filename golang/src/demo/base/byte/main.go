package main

import "fmt"

func main() {
	// byte字符 属于int类型,单引号表示
	var b = 'a'
	fmt.Printf("%v \n", b)         // 97
	fmt.Printf("%c -- %[1]T\n", b) // 'a'

	var s = "this"
	fmt.Printf("%v -- %c --%[1]T \n", s, s[2]) // %c 原样输出字符

	var z = "李"
	fmt.Printf("%v %[1]T \n", z)

	var str = "中文 golang"
	for i := 0; i < len(str); i++ { // byte 类型
		fmt.Printf("%v(%[1]c - %[1]T) \n", str[i])
	}
	// 循环字符串,如果字符串中存在中文则使用reage循环遍历（推荐使用）
	for _, v := range str { // rune 类型
		fmt.Printf("%v(%[1]c - %[1]T) \n", v)
	}

	// 修改字符串中的字符
	var str1 = "golang php"
	byteStr := []byte(str1) // 强制类型转换
	byteStr[6] = ','
	str1 = string(byteStr)
	fmt.Println(str1)

	runeStr := []rune(str1) // 强制类型转换
	runeStr[6] = ','
	str1 = string(runeStr)
	fmt.Println(str1)
}
