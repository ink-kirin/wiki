package main

import (
	"encoding/json"
	"fmt"
)

type Persion struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func main() {
	var s = Persion{
		Id:   11,
		Name: "ss",
		Age:  30,
		Sex:  "男",
	}
	fmt.Printf("%#v \n", s) // main.Persion{Id:11, Name:"ss", Age:30, Sex:"男"}
	jb, _ := json.Marshal(s)
	js := string(jb)
	fmt.Println(js) // {"id":11,"name":"ss","age":30,"sex":"男"}
}
