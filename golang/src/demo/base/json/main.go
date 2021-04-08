package main

import (
	"encoding/json"
	"fmt"
)

// Host 定义数据类型
type Host struct {
	IP   string
	Name string
	age  int // 私有属性不能被JSON包访问
}

func main() {
	b := []byte(`{"IP":"192.168.12.194","Name":"secket"}`)
	m := Host{}

	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}

	fmt.Println("m:", m)
	fmt.Println("m.IP:", m.IP)
	fmt.Println("m.Name:", m.Name)

	// 结构体转JSON
	var s = Host{
		IP:   "182.11.23.12",
		Name: "jia",
		age:  22,
	}
	jb, _ := json.Marshal(s) // return byte[]
	fmt.Println(string(jb))

	// JSON转结构体
	var t = `{"IP":"182.11.23.12","Name":"jia","age":23}`
	var h Host
	error := json.Unmarshal([]byte(t), &h)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(h)
}
