package main

import "fmt"

func main() {
	// 数组的长度是类型的一部分
	// 数组的初始化 int默认0 string默认空
	var arr1 [3]int
	var arr2 [4]int
	var strArr [3]string
	fmt.Printf("%T, %T, %T \n", arr1, arr2, strArr) // [3]int, [4]int, [3]string
	fmt.Printf("%v, %v, %v \n", arr1, arr2, strArr) // [0 0 0], [0 0 0 0], [  ]

	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println(arr1) // [1 2 3]

	var arr3 = [3]int{1, 2, 3}
	fmt.Println(arr3) // [1 2 3]

	var arr4 = [3]string{"php", "golang", "java"}
	fmt.Println(arr4) // [php golang java]

	// 数组长度值不固定
	var arr5 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr5, len(arr5)) // [1 2 3 4 5] 5

	// 修改数组的值
	var arr6 = [...]string{"php", "golang", "java"}
	arr6[2] = "Java"
	fmt.Println(arr6) // [php golang Java]

	// 初始化定义下标，未赋值的下标值为0
	var arr7 = [...]int{1: 2, 4: 8, 5: 9}
	fmt.Println(arr7, len(arr7)) // [0 2 0 0 8 9] 6

	// 循环遍历数组 for  for range
	for k, v := range arr7 {
		fmt.Println(k, v)
	}

	// 计算数组中的最大值，并获取下标
	var arr8 = [...]int{4, 32, 13, 0, 89}
	var max = arr8[0]
	var index = 0
	for i := 0; i < len(arr8); i++ {
		if max < arr8[i] {
			max = arr8[i]
			index = i
		}
	}
	fmt.Printf("最大值：%d, 索引：%d \n", max, index)

	// 从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标
	var arr9 = [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(arr9); i++ {
		for j := i + 1; j < len(arr9); j++ {
			if arr9[i]+arr9[j] == 8 {
				fmt.Printf("(%d, %d) \n", i, j)
			}
		}
	}

	// 数组是值类型, 切片是引用类型

	// 数组
	var q = [...]int{1, 2, 3} // 一位数组
	var a = q
	q[0] = 8
	fmt.Println(q, a) // [8 2 3] [1 2 3]

	// 切片
	var q1 = []int{1, 2, 3}
	var a1 = q
	q1[0] = 8
	a1[2] = 10
	fmt.Println(q1, a1) // [8 2 3] [8 2 10]

	// 多维数组
	var e = [3][2]string{
		{"北京", "上海"},
		{"石家庄", "天津"},
		{"深圳", "郑州"},
	}
	fmt.Println(e) // [[北京 上海] [石家庄 天津] [深圳 郑州]]
	// 遍历多维数组
	for _, v1 := range e {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	// 长度不固定的多维数组(第二层开始不支持... 自动推导)
	var e1 = [...][2]string{
		{"北京", "上海"},
		{"石家庄", "天津"},
		{"深圳", "郑州"},
	}
	fmt.Println(e1)

	// 声明切片
	var s = []int{1, 2, 3}
	fmt.Println(s)

	var arr10 = [5]int{1, 3, 5, 7, 8} // array
	s1 := arr10[1:4]                  // slice 包含1，不包含4
	fmt.Printf("%v, %T \n", s1, s1)

	// make 声明切片
	var m = make([]int, 4, 8)
	fmt.Println(m) // [0 0 0 0]
	m[0] = 22
	m[1] = 4
	m[2] = 90
	m[3] = 64
	fmt.Println(m) // [22 4 90 64]

	// 切片扩容 append
	m = append(m, 23)
	fmt.Println(m) // [22 4 90 64 23]

	// 合并切片
	m = append(m, s...)
	fmt.Println(m) // [22 4 90 64 23 1 2 3]

	// copy 复制切片
	s2 := make([]int, 3, 3)
	copy(s2, s) // copy(切片，被复制切片)
	fmt.Println(s2)

	// 删除切片的值
	var q2 = []int{1, 2, 3, 4, 5, 6}
	q2 = append(q2[:2], q2[3:]...)
	fmt.Println(q2) // [1 2 4 5 6]
}
