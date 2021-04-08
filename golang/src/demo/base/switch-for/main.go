package main

import "fmt"

func main() {
	var ext = ".html"
	switch ext {
	case ".php":
		fmt.Printf("text/php")
		break
	case ".css":
		fmt.Printf("text/css")
		break
	default:
		fmt.Printf("text")
		break
	}

	var n = 5
	switch n {
	case 1, 3, 5, 7, 9:
		fmt.Printf("奇数")
		break
	case 2, 4, 6, 8, 10:
		fmt.Printf("偶数")
		break
	}

	// 如果判断条件是表达式则switch不需要传值
	var age = 20
	switch {
	case age < 22:
		fmt.Printf("未成年")
		break
	case age >= 22 && age <= 48:
		fmt.Printf("青年")
		break
	case age > 48:
		fmt.Printf("老年")
		// fallthrough // 穿透下层
		break
	default:
		fmt.Printf("参数错误")
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
