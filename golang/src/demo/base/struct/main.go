package main

import "fmt"

/*
	首字母大写表示是共有的，在其他包里可以使用
*/

// Info 嵌套结构体 继承
type Info struct {
	Data string
	Persion
}

type Animal struct {
	err string
	*Persion
}

// Persion 定义结构体
type Persion struct {
	name  string
	age   int
	sex   string
	Hobby []string
	Ind   map[string]string
}

// 值传递
func (p Persion) info(name string, age int) {
	p.name = name
	p.age = age
}

// 引用传递
func (p *Persion) SetInfo(name string, age int) {
	p.name = name
	p.age = age
}

// 结构体
func main() {
	var an = Animal{
		err: "error",
		Persion: &Persion{
			name: "123",
		},
	}
	fmt.Printf("%#v %v \n", an, an.name)

	var i Info
	i.Data = "data"
	i.name = "simba"
	fmt.Println(i)

	var per Persion
	per.name = "er"
	per.age = 22
	per.Hobby = make([]string, 3)
	per.Hobby[0] = "a"
	per.Hobby[1] = "b"
	per.Hobby[2] = "c"
	per.Ind = make(map[string]string)
	per.Ind["address"] = "ddd"
	per.Ind["call"] = "cccc"
	fmt.Println(per)

	var a = Persion{
		name: "lisi",
		age:  30,
	}
	fmt.Println(a)
	a.SetInfo("simba", 29)
	fmt.Println(a)
	a.info("ss", 22)
	fmt.Println(a)

	// 实例化结构体
	// 方法一
	var p Persion // 实例化Persion结构体
	p.name = "simba"
	p.age = 30
	p.sex = "男"
	fmt.Printf("%#v \n", p) // main.Persion{name:"simba", age:30, sex:"男"}
	// 方法二
	var p1 = new(Persion) // 指针类型
	p1.name = "李"
	p1.age = 20
	p1.sex = "男"
	fmt.Printf("%#v \n", p1) // &main.Persion{name:"李", age:20, sex:"男"}
	(*p1).name = "王"         // 修改指针的值
	fmt.Printf("%#v \n", p1) // &main.Persion{name:"王", age:20, sex:"男"}
	// 方法三
	var p2 = &Persion{}
	p2.name = "simba"
	p2.age = 29
	p2.sex = "男"
	fmt.Printf("%#v \n", p2) // &main.Persion{name:"simba", age:29, sex:"男"}
	// 方法四 (推荐)
	var p3 = Persion{
		name: "lsi",
		age:  28,
		sex:  "男",
	}
	fmt.Printf("%#v \n", p3) // main.Persion{name:"lsi", age:28, sex:"男"}
	// 方法五 （推荐）
	var p4 = &Persion{
		name: "ali",
		age:  28,
		sex:  "男",
	}
	fmt.Printf("%#v \n", p4) // &main.Persion{name:"ali", age:28, sex:"男"}
	// 方法六
	var p5 = &Persion{
		name: "ali",
	}
	fmt.Printf("%#v \n", p5) // &main.Persion{name:"ali", age:0, sex:""}
	// 方法七
	// var p6 = &Persion{ // 如果省略key的话，必须全部字段都赋值
	// 	"nuw",
	// 	28,
	// 	"男",
	// }
	// fmt.Printf("%#v \n", p6) // &main.Persion{name:"nuw", age:28, sex:"男"}
}
