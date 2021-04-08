package main

import (
	"fmt"
	"sort"
)

// 冒泡排序 升序
func asc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				temp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = temp
			}
		}
	}
	return slice
}

// 冒泡排序 降序 // 没有返回值，修改传入的参数值
func desc(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] < slice[j+1] {
				temp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = temp
			}
		}
	}
}

func main() {
	// sort 升序
	var list1 = []int{6, 4, 7, 9, 5, 8}
	var list2 = []float64{3.4, 2.3, 0.5, 9.3, 7, 1}
	var list3 = []string{"a", "va", "bg", "re"}

	sort.Ints(list1)
	sort.Float64s(list2)
	sort.Strings(list3)

	fmt.Println(list1)
	fmt.Println(list2)
	fmt.Println(list3)

	// sort 降序
	sort.Sort(sort.Reverse(sort.IntSlice(list1)))
	sort.Sort(sort.Reverse(sort.Float64Slice(list2)))
	sort.Sort(sort.Reverse(sort.StringSlice(list3)))

	fmt.Println(list1)
	fmt.Println(list2)
	fmt.Println(list3)

	var list4 = []int{6, 4, 7, 9, 5, 8}
	fmt.Println(asc(list4))

	var list5 = []int{6, 4, 7, 9, 5, 8}
	desc(list5)
	fmt.Println(list5)

}
