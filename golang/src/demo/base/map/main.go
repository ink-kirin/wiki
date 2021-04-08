package main

import "fmt"

func main() {
	// 创建map
	var info = make(map[string]string) // map类型必须使用make分配内存空间才可以使用
	info["name"] = "simba"
	fmt.Println(info)

	// 初始化map
	var user = map[string]string{
		"age": "20", // 必须,结尾
	}
	fmt.Println(user)

	// 循环遍历map
	var userinfo = map[string]string{
		"name":   "simba",
		"age":    "30",
		"sex":    "男",
		"height": "180cm",
	}
	for k, val := range userinfo {
		fmt.Println(k, val)
	}

	// map curd
	fmt.Println(userinfo["name"]) // 获取
	v, ok := userinfo["age"]      // 查找
	fmt.Println(v, ok)            // 30 true
	delete(userinfo, "height")    // 删除
	fmt.Println(userinfo)         // map[age:30 name:simba sex:男]
}
