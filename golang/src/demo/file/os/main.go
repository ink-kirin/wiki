package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// 读取文件
func read(url string) string {
	// 打开文件
	file, err := os.Open(url)
	defer file.Close()
	if err != nil {
		return err.Error()
	}
	// 读取文件
	var strSlice []byte
	var tempSlice = make([]byte, 128)
	for {
		n, err := file.Read(tempSlice) // 读取
		if err == io.EOF {
			break
		}
		if err != nil {
			return "读取失败:" + err.Error()
		}
		strSlice = append(strSlice, tempSlice[:n]...)
	}
	return string(strSlice)
}

// 写入文件
func writer(url string, str string) string {
	file, err := os.OpenFile(url, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		return err.Error()
	}
	// 写入文件
	for i := 0; i < 10; i++ {
		file.WriteString("os写入字符串:" + str + strconv.Itoa(i) + "\r\n")
		file.Write([]byte("os写入byte类型:" + str + strconv.Itoa(i) + "\r\n"))
	}
	return "Success"
}

// 通过文件流方式拷贝文件
func copy(src string, dst string) (err error) {
	s, err := os.Open(src)
	d, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, 0666)
	defer s.Close()
	defer d.Close()
	if err != nil {
		return err
	}
	var t = make([]byte, 1024)
	for {
		n, err := s.Read(t)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if _, err := d.Write(t[:n]); err != nil {
			return err
		}
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

	s1 := "/Users/simba/test.log"
	s2 := "/Users/simba/test2.log"
	err := copy(s1, s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("拷贝完成")

	// 创建目录
	os.Mkdir("/Users/simba/test/", os.ModePerm) // os.ModePerm 0777

	// 创建多级目录
	os.MkdirAll("/Users/simba/test/1/", os.ModePerm) // os.ModePerm 0777

	// 删除目录
	os.Remove("/Users/simba/test/2")
	// 删除文件
	os.Remove("/Users/simba/test/test.log")
	// 删除目录下所有文件
	os.RemoveAll("/Users/simba/test/2")

	// 重命名文件
	err = os.Rename("/Users/simba/test.log", "/Users/simba/test4.log")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("重命名成功")
}
