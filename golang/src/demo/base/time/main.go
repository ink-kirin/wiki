package main

import (
	"fmt"
	"time"
)

func newTime() string {
	timeObj := time.Now()      // 当前时间对象
	year := timeObj.Year()     // 年
	month := timeObj.Month()   // 月
	day := timeObj.Day()       // 日
	hour := timeObj.Hour()     // 时
	minute := timeObj.Minute() // 分
	second := timeObj.Second() // 秒
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
	// 注：%02中的2表示宽度，如果整数不够2位就补位0
}

func main() {
	fmt.Printf("%s -- %[1]T \n", newTime())

	// 使用time内置方法
	/*
		2006 年
		01 月
		02 日
		03 时 	// 03表示12小时制  15表示24小时制
		04 分
		05 秒
	*/
	time1 := time.Now()
	var str = time1.Format("2006-01-02 15:04:05") // 格式化时间
	fmt.Println(str)

	// 获取当前(秒)时间戳
	unix := time1.Unix()
	fmt.Println(unix)
	// 获取当前(毫秒)时间戳
	unixtime := time1.UnixNano() / 1e6
	fmt.Println(unixtime)
	// 获取当前(纳秒)时间戳
	nanotime := time1.UnixNano()
	fmt.Println(nanotime)
	// 纳秒转换为秒
	fmt.Println(time.Now().UnixNano() / 1e9)

	// 时间戳转日期
	var str1 int64 = 1592408980
	time2 := time.Unix(str1, 0)
	var str2 = time2.Format("2006-01-02 15:04:05")
	fmt.Println(str2)

	// 日期转时间戳
	var str3 = "2020-06-17 23:49:40"
	time3, _ := time.ParseInLocation("2006-01-02 15:04:05", str3, time.Local)
	fmt.Println(time3.Unix())

	// time.Millisecond // 1毫秒
	// time.Second      // 1秒

	// time.Now().Add(time.Hour) // 当前时间增加一小时

	/*
	 定时器
	*/
	// 方法一：time.NewTicker 生成定时器
	ticker := time.NewTicker(time.Second) // 初始化定时器，间隔时间1秒
	n := 5
	for v := range ticker.C {
		fmt.Println(v)
		n--
		if n == 0 {
			ticker.Stop() // 终止定时器
			break
		}
	}

	// time.Sleep 休眠
	time.Sleep(time.Second * 5)
	fmt.Println("休眠5秒")
}
