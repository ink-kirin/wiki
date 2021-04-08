package main

import "fmt"

// 接口是一个规范，是一种类型
type Usber interface {
	start()
	stop()
}

// 电脑
type Computer struct {
}

func (c Computer) work(usb Usber) {
	usb.start()
	usb.stop()
}

// 手机
type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "start")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "stop")
}

// 相机
type Camera struct {
}

func (p Camera) start() {
	fmt.Println("start")
}

func (p Camera) stop() {
	fmt.Println("stop")
}

func (p Camera) run() {
	fmt.Println("run")
}

// 定义接口的标准
func main() {
	var p Usber
	p = Phone{
		Name: "小米",
	}
	p.start()

	var c Camera
	var c1 Usber = c
	c1.stop()
	c.run()

	var computer = Computer{}
	var phone = Phone{
		Name: "小米",
	}
	var camera = Camera{}
	computer.work(phone)
	computer.work(camera)
}
