package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名:%v 年龄:%v 成绩:%v", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) Print() string {
	fmt.Println("打印....")
	return "Print 完成"
}

func PrintStruct(s interface{}) {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("参数不是结构体")
		return
	}
	// Field 获取结构体的字段
	field0 := t.Field(0)
	fmt.Println("名称:", field0.Name)
	fmt.Println("类型:", field0.Type)
	fmt.Println("Tag:", field0.Tag.Get("json"))

	// FieldByName 获取结构体字段
	field1, ok := t.FieldByName("Age")
	if ok {
		fmt.Println("名称:", field1.Name)
		fmt.Println("类型:", field1.Type)
		fmt.Println("Tag:", field1.Tag.Get("json"))
	}

	// NumField 获取结构体有几个字段
	num := t.NumField()
	fmt.Println("字段数量:", num)

	// 获取结构体的值
	v := reflect.ValueOf(s)
	fmt.Println(v.FieldByName("Name"))

}

func fn(s interface{}) {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("参数不是结构体")
		return
	}

	// Method 获取结构体的方法
	m := t.Method(0) // 结构体的方法顺序是按照ASCII排序的
	fmt.Println(m.Name)
	fmt.Println(m.Type)

	// MethodByName 获取结构体的方法
	m1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println(m1.Name)
		fmt.Println(m1.Type)
	}

	// 执行结构体的方法
	v := reflect.ValueOf(s)
	p := v.MethodByName("Print").Call(nil) // 返回类型是切片
	fmt.Println(p[0])

	// 执行方法传递参数
	var params []reflect.Value // reflect.Value 类型的切片
	params = append(params, reflect.ValueOf("kirin"))
	params = append(params, reflect.ValueOf(30))
	params = append(params, reflect.ValueOf(99))
	v.MethodByName("SetInfo").Call(params) // 执行方法传入参数

	i := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(i)

	// 获取方法数量
	fmt.Println("方法数量:", t.NumMethod())
}

// 修改结构体属性的值
func fnChange(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Elem().Kind() != reflect.Struct {
		fmt.Println("参数不是结构体指针类型")
		return
	} else if t.Kind() != reflect.Ptr {
		fmt.Println("参数不是结构体指针类型")
		return
	}
	// 修改结构体属性的值
	name := v.Elem().FieldByName("Name")
	name.SetString("啦啦啦")

	age := v.Elem().FieldByName("Age")
	age.SetInt(21)
}

func main() {
	stu := Student{
		Name:  "simba",
		Age:   23,
		Score: 20,
	}
	PrintStruct(stu)
	fmt.Println("fn--------")
	fn(&stu)

	// 修改结构体的属性
	fnChange(&stu)
	fmt.Println("修改后的值:", stu)
}
