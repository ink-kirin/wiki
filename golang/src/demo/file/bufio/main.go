package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// 读取文件
func read(url string) string {
	file, err := os.Open(url)
	defer file.Close()
	if err != nil {
		return err.Error()
	}
	// bufio 读取文件
	var fileStr string
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 表示一次读取一行
		if err == io.EOF {
			fileStr += str
			break
		}
		if err != nil {
			return err.Error()
		}
		fileStr += str
	}
	return fileStr
}

// 写入文件
func writer(url string, str string) string {
	file, err := os.OpenFile(url, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		return err.Error()
	}
	// 写入文件
	w := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		w.WriteString("bufio写入字符串:" + str + strconv.Itoa(i) + "\r\n")
		w.Write([]byte("bufio写入byte类型:" + str + strconv.Itoa(i) + "\r\n"))
	}
	w.Flush()
	return "Success"
}

func main() {
	url := "/Volumes/workspace/video/mongo_shell.txt"
	r := read(url)
	fmt.Println(r)

	wurl := "/Users/simba/test.log"
	s := "Sibma"
	w := writer(wurl, s)
	fmt.Println(w)
}
