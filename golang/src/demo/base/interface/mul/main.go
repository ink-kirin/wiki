package main

import "fmt"

type Animlaer interface {
	SetName(string)
	GetName() string
}

type Log interface {
	Add(string) string
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d Dog) GetName() string {
	return d.Name
}

func (d *Dog) Add(name string) string {
	d.Name = name
	return name
}

// 一个结构体实现多个方法
func main() {
	var a Animlaer = &Dog{
		Name: "haha",
	}
	fmt.Println(a.GetName())
	a.SetName("lili")
	fmt.Println(a.GetName())

	var l Log = &Dog{
		Name: "1010",
	}
	v := l.Add("3030")
	fmt.Println(v)
}
