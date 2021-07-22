package main

import (
	"fmt"
	"strconv"
)

// 整数转换二进制
func convertToBin(n int) string {
	if n == 0 {
		return "00"
	}
	res := ""
	for ; n > 0; n /= 2 {
		l := n % 2
		res = strconv.Itoa(l) + res
	}
	return res
}

func main() {

	fmt.Println(
		convertToBin(5),
		convertToBin(17),
	)

	var ext = ".html"
	switch ext {
	case ".php":
		fmt.Println("text/php")
	case ".css":
		fmt.Println("text/css")
	default:
		fmt.Println("text")
	}

	var n = 5
	switch n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	}

	// 如果判断条件是表达式则switch不需要传值
	var age = 32
	switch {
	case age < 22:
		fmt.Println("未成年")
	case age >= 22 && age <= 48:
		fmt.Println("青年")
	case age > 48:
		fmt.Println("老年")
		// fallthrough // 穿透下层
	default:
		fmt.Println("参数错误")
	}

	for i := 1; i < 10; i++ {
		fmt.Println(i)
		if i == 4 {
			break
		}
	}
	/*
		break 跳出循环，执行后面的程序
		continue 跳出本地循环，继续下次循环
	*/
	// 对于for/select /switch ,Label必须紧挨着他们
End: // label
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		if i == 4 {
			break End
		}
	}
	fmt.Printf("End")
}
