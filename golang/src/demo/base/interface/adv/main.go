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
	// 判断sub的类型
	if _, ok := usb.(Phone); ok { // 类型断言
		usb.start()
	} else {
		usb.stop()
	}
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

// 一个接口多个结构体
func main() {
	var computer = Computer{}
	var phone = Phone{
		Name: "小米",
	}
	var camera = Camera{}
	computer.work(phone)
	computer.work(camera)
}
