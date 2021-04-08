package main

import "fmt"

// 接口是一个规范，是一种类型
type Usber interface {
	start()
	stop()
}

// 手机
type Phone struct {
	Name string
}

func (p Phone) start() { // 值类型接受者
	fmt.Println(p.Name, "start")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "stop")
}

type Camera struct {
}

func (p *Camera) start() { // 指针类型接受者
	fmt.Println("start")
}

func (p *Camera) stop() {
	fmt.Println("stop")
}

// 指针类型
func main() {
	var p Usber
	p = Phone{
		Name: "小米",
	}
	p.start()

	var p1 = &Camera{}
	var u Usber = p1
	u.start()
}
