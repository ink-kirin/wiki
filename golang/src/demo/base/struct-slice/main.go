package main

import (
	"encoding/json"
	"fmt"
)

type Students struct {
	Id   int
	Name string
	Age  int
}

type Class struct {
	Title    string
	Students []Students
}

func main() {
	var c = Class{
		Title:    "001班",
		Students: make([]Students, 0), // 实例化切片
	}
	for i := 0; i < 10; i++ {
		var s = Students{
			Id:   i,
			Name: fmt.Sprintf("stu_%v", i),
			Age:  2 + i,
		}
		c.Students = append(c.Students, s)
	}
	fmt.Println(c) // {001班 [{0 stu_0 2} {1 stu_1 3} {2 stu_2 4} {3 stu_3 5} {4 stu_4 6} {5 stu_5 7} {6 stu_6 8} {7 stu_7 9} {8 stu_8 10} {9 stu_9 11}]}
	jb, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	} else {
		js := string(jb)
		fmt.Println(js)
	}

	var str = `{"Title":"001班","Students":[{"Id":0,"Name":"stu_0","Age":2},{"Id":1,"Name":"stu_1","Age":3},{"Id":2,"Name":"stu_2","Age":4},{"Id":3,"Name":"stu_3","Age":5},{"Id":4,"Name":"stu_4","Age":6},{"Id":5,"Name":"stu_5","Age":7},{"Id":6,"Name":"stu_6","Age":8},{"Id":7,"Name":"stu_7","Age":9},{"Id":8,"Name":"stu_8","Age":10},{"Id":9,"Name":"stu_9","Age":11}]}`
	cl := &Class{}
	err = json.Unmarshal([]byte(str), cl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v \n", cl)

}
