package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// 读取接口
type Reader interface {
	Read(rc chan []byte)
}

// 写入接口
type Writer interface {
	Write(wc chan string)
}

type ReadFromFile struct {
	path string
}

type WriteToInfluxDB struct {
	influxDBDsn string
}

type LogProcess struct {
	rc    chan []byte
	wc    chan string
	read  Reader
	write Writer
}

func (r *ReadFromFile) Read(rc chan []byte) {
	// 打开文件
	l, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file err:%s", err.Error()))
	}
	// 从文件末尾开始逐行读取
	l.Seek(0, 2)
	rd := bufio.NewReader(l)
	for {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes err:%s", err.Error()))
		}
		rc <- line[:len(line)-1]
	}
}

func (r *WriteToInfluxDB) Write(wc chan string) {
	for v := range wc {
		fmt.Println(v)
	}
}

// 解析
func (l *LogProcess) Process() {
	for v := range l.rc {
		l.wc <- strings.ToUpper(string(v))
	}
}

func main() {
	r := &ReadFromFile{
		path: "error.log",
	}
	w := &WriteToInfluxDB{
		influxDBDsn: "username",
	}
	lp := &LogProcess{
		rc:    make(chan []byte),
		wc:    make(chan string),
		read:  r,
		write: w,
	}

	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(20 * time.Second)
}
