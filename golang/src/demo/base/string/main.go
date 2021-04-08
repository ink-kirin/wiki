package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "hello, world"
	fmt.Println(len(s))     // 12
	fmt.Println(s[0], s[7]) // 104 119 ('h' and 'w')
	fmt.Println(s[:5])      // "hello"
	fmt.Println(s[7:])      // "world"
	fmt.Println(s[:])       // "hello, world"
	// 字符串拼接
	fmt.Println("goodbye" + s[5:]) // "goodbye, world"

	s = "left foot"
	t := s
	s += ", right root"
	fmt.Println(s)                          // "left root, right root"
	fmt.Println(t)                          // "left root"
	fmt.Println(fmt.Sprintf("%s %s", s, t)) // fmt.Sprint() 拼接字符串
	// 因为字符串是不可修改的，因此尝试修改字符串内部数据的操作也是被禁止的
	// s[0] = 'L' // compile error: cannot assign to s[0]
	/*
		ASCII控制代码:
			\a 响铃
			\b 退格
			\f 换页
			\n 换行
			\r 回车
			\t 制表符
			\v 垂直制表符
			\' 单引号 (只用在 '\'' 形式的rune符号面值中)
			\" 双引号 (只用在 "..." 形式的字符串面值中)
			\\ 反斜杠
	*/
	f := `this is 1
this is 2
this is 3
`
	fmt.Println(f)
	fmt.Println(len(f)) // len() 获取字符串长度

	// strings.Split 分割字符串 	return 切片
	var str = "123-234-453"
	arr := strings.Split(str, "-") // [123 234 453]
	fmt.Println(arr)

	// strings.Join() list 转 string    return string
	str1 := strings.Join(arr, "-") // 123-234-453
	fmt.Println(str1)

	arr1 := []string{"php", "golang"} // 定义切片
	fmt.Println(strings.Join(arr1, ","))

	// strings.Contains 判断是否包含
	str2 := "this is golang"
	str3 := "golang"
	fmt.Println(strings.Contains(str2, str3)) // str2是否包含str3

	// strings.HasPrefix 前缀   strings.hasSuffix 后缀
	fmt.Println(strings.HasPrefix(str2, str3))

	// strings.Index 判断子字符串或字符在父字符串中出现的位置（索引）,查找到返回下标,否则返回-1, 下标从0开始
	fmt.Println(strings.Index(str2, "l"))

	// strings.LastIndex 最后出现位置的索引下标
	fmt.Println(strings.LastIndex(str2, "s"))

	var v1 int = 20
	var v2 float64 = 13.3321
	var v3 bool = false
	var v4 byte = 'a'
	fmt.Println(fmt.Sprintf("%d", v1)) // int类型转换string
	fmt.Println(fmt.Sprintf("%f", v2)) // float转换string
	fmt.Println(fmt.Sprintf("%t", v3)) // bool转换string
	fmt.Println(fmt.Sprintf("%c", v4)) // byte转换string
	/*
		int - %d
		float - %f
		bool - %t
		byte - %c
	*/
	strconv.FormatInt(int64(v1), 10)     // 参数1：必须是int64，参数2：进制，10表示10进制
	strconv.FormatFloat(v2, 'f', -1, 64) // 参数1:要转换的值，参数2:格式化类型'f','b','e'，参数3:保留的小数点(-1不格式化小数点)，参数4:格式化的类型 64|32
	strconv.FormatBool(v3)
	strconv.FormatUint(uint64(v4), 10) // 参数1：必须是uint64，参数2：进制，10表示10进制

	// string 转 int
	var s1 = "1234"
	num1, _ := strconv.ParseInt(s1, 10, 64) // ParseInt 参数1：string，参数2：进制，参数3：位数 32 64  return int error
	fmt.Println(num1)
	// ParseFloat string -> float

	// 遍历中文字符串
	str = "我是中文字符串"         // 一个中文占3个字符
	for k, v := range str { // range retrun key=>value
		fmt.Printf("key=%d value=%c \n", k, v)
	}
}
