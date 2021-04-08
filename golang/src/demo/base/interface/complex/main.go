package main

import "fmt"

type Animlaer interface {
	SetName(string)
	GetName() string
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

func main() {
	var a Animlaer = &Dog{
		Name: "haha",
	}
	fmt.Println(a.GetName())
	a.SetName("lili")
	fmt.Println(a.GetName())
}
