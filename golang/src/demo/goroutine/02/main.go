package main

import (
	"fmt"
	"time"
)

func main() {
	// 计算1-120000之间的素数
	start := time.Now().UnixNano() / 1e6 // 毫秒
	for num := 2; num < 120000; num++ {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			// fmt.Println(num, "是素数")
		}
	}
	end := time.Now().UnixNano() / 1e6 // 毫秒
	fmt.Println(end-start, "毫秒")
}
