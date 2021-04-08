package main

import (
	"fmt"
	"io/ioutil"
)

// 读取文件
func read(url string) string {
	byteStr, err := ioutil.ReadFile(url)
	if err != nil {
		return err.Error()
	}
	return string(byteStr)
}

// 写入文件
func writer(url string, str string) string {
	err := ioutil.WriteFile(url, []byte(str), 0666) // 清空文件写入内容
	if err != nil {
		return err.Error()
	}
	return "Success"
}

// 复制文件
func copy(src string, dst string) (err error) {
	byteStr, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dst, byteStr, 0666)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	url := "/Volumes/workspace/video/mongo_shell.txt"
	r := read(url)
	fmt.Println(r)

	wurl := "/Users/simba/test.log"
	s := "Sibma"
	w := writer(wurl, s)
	fmt.Println(w)

	// 复制文件
	u2 := "/Users/simba/test1.log"
	err := copy(wurl, u2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("复制完成")
}
