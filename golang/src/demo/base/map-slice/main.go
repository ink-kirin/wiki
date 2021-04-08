package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// 元素为map的切片 map[]不初始化的默认值为nil
	var userinfo = make([]map[string]string, 2)
	if userinfo[0] == nil {
		userinfo[0] = make(map[string]string)
		userinfo[0]["name"] = "simba"
		userinfo[0]["age"] = "20"
		userinfo[0]["sex"] = "男"
	}
	if userinfo[1] == nil {
		userinfo[1] = make(map[string]string)
		userinfo[1]["name"] = "lai"
		userinfo[1]["age"] = "18"
		userinfo[1]["sex"] = "女"
	}
	fmt.Println(userinfo)
	// 循环遍历
	for _, v := range userinfo {
		for k, val := range v {
			fmt.Println(k, val)
		}
	}

	// 值为切片的map
	var info = make(map[string][]string) // []string 切片类型是string
	info["hobby"] = []string{
		"吃饭",
		"睡觉",
	}
	info["work"] = []string{
		"PHP",
		"Java",
		"Golang",
	}
	fmt.Println(info)
	for k, v := range info {
		// fmt.Println(k, v)
		for key, val := range v {
			fmt.Println(k, key, val)
		}
	}

	// map 是引用类型
	var user = make(map[string]string)
	user["name"] = "lsi"
	user1 := user

	user1["name"] = "simba"
	fmt.Println(user, user1)

	// map 排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[6] = 46
	map1[3] = 35

	// key 升序
	var keySlice []int
	for k := range map1 { // for k, _ := range map1 如果需要忽略第二个值，则可以省略不写
		keySlice = append(keySlice, k)
	}
	fmt.Println(keySlice)
	sort.Ints(keySlice)
	for _, val := range keySlice {
		fmt.Printf("key=%d value=%d \n", val, map1[val])
	}

	// 统计单词在字符串中出现的次数
	var str = "how do you do"
	var strSlice = strings.Split(str, " ")
	var strMap = make(map[string]int)
	for _, v := range strSlice {
		strMap[v]++
	}
	fmt.Println(strMap)
}
