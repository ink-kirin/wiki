package main

import "fmt"

func main() {
	// [9,8,7,6,5,4]倒序  选择排序
	var str = []int{9, 8, 7, 6, 5, 4}
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] > str[j] {
				temp := str[i]
				str[i] = str[j]
				str[j] = temp
			}
		}
	}
	fmt.Println(str)

	// 冒泡排序
	var s = []int{9, 8, 7, 6, 5, 4}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				temp := s[j]
				s[j] = s[j+1]
				s[j+1] = temp
			}
		}
	}
	fmt.Println(s)
}
