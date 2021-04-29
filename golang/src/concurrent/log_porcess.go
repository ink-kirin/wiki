package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// 读取接口
type Reader interface {
	Read(rc chan []byte)
}

// 写入接口
type Writer interface {
	Write(wc chan *Message)
}

type LogProcess struct {
	rc    chan []byte
	wc    chan *Message
	read  Reader
	write Writer
}

type ReadFromFile struct {
	path string
}

type WriteToInfluxDB struct {
	influxDBDsn string
}

// 日志参数
type Message struct {
	TimeLocal                    time.Time
	BytesSent                    int
	Path, Method, Scheme, Status string
	UpstreamTime, RequestTime    float64
}

// 系统状态监控
type SystemInfo struct {
	HandleLine   int     `json:"handleline"`   // 总处理日志行数
	Tps          float64 `json:"tps"`          // 系统吞吐量
	ReadChanLen  int     `json:"readChanLen"`  // read channel 长度
	WriteChanLen int     `json:"writeChanLen"` // write channel 长度
	RunTime      string  `json:"runTime"`      // 运行总时间
	ErrNum       int     `json:"errNum"`       // 错误数
}

const (
	TypeHandleLine = 0
	TypeErrNum     = 1
)

var TypeMonitorChan = make(chan int, 200)

type Monitor struct {
	startTime time.Time
	data      SystemInfo
	tpsSli    []int
}

func (m *Monitor) start(lp *LogProcess) {

	go func() {
		for n := range TypeMonitorChan {
			switch n {
			case TypeErrNum:
				m.data.ErrNum += 1
			case TypeHandleLine:
				m.data.HandleLine += 1
			}
		}
	}()

	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for {
			<-ticker.C
			m.tpsSli = append(m.tpsSli, m.data.HandleLine)
			if len(m.tpsSli) > 2 {
				m.tpsSli = m.tpsSli[1:]
			}
		}
	}()

	http.HandleFunc("/monitor", func(writer http.ResponseWriter, request *http.Request) {
		m.data.RunTime = time.Since(m.startTime).String()
		m.data.ReadChanLen = len(lp.rc)
		m.data.WriteChanLen = len(lp.wc)
		if len(m.tpsSli) >= 2 {
			m.data.Tps = float64(m.tpsSli[1]-m.tpsSli[0]) / 5
		}
		ret, _ := json.MarshalIndent(m.data, "", "\t")
		io.WriteString(writer, string(ret))
	})

	http.ListenAndServe(":9193", nil)
}

// 打开
func (r *ReadFromFile) Read(rc chan []byte) {
	// 打开文件
	l, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file err:%s", err.Error()))
	}
	// 从文件末尾开始逐行读取
	_, err = l.Seek(0, 2) // 偏移量移至末尾
	if err != nil {
		panic(fmt.Sprintf("Seek err: %s", err.Error()))
	}

	rd := bufio.NewReader(l)
	for {
		// 需要优化分割日志文件读取
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes err:%s", err.Error()))
		}
		TypeMonitorChan <- TypeHandleLine
		rc <- line[:len(line)-1]
	}
}

// 写入
func (r *WriteToInfluxDB) Write(wc chan *Message) {
	// Create a new client using an InfluxDB server base URL and an authentication token
	client := influxdb2.NewClient("http://localhost:8086", "xcQhIN_le_PbmjljLzX_F7ok-37HPe2w6QYN6RmX9KMrij7Y0fjjf0oom5JwVZy5djrJEAxVj8wt_7p9CXjXqA==")

	for v := range wc {
		// Use blocking write client for writes to desired bucket
		writeAPI := client.WriteAPIBlocking("simba", "dev")
		tags := map[string]string{
			"Path":   v.Path,
			"Method": v.Method,
			"Scheme": v.Scheme,
			"Status": v.Status,
		}
		fields := map[string]interface{}{
			"UpstreamTime": v.UpstreamTime,
			"RequestTime":  v.RequestTime,
			"BytesSent":    v.BytesSent,
		}
		// Create point using full params constructor
		p := influxdb2.NewPoint("nginx_log",
			tags,
			fields,
			v.TimeLocal)
		// write point immediately
		err := writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Write Success")
	}
}

// 解析
func (l *LogProcess) Process() {

	// r := regexp.MustCompile("\\[([\\w-:\\+\\.]*)\\][\\s](htt[ps]://[\\w./\\?=&]+)[\\s]([0-9]+)[\\s]([\\w]+)[\\s]({.+})[\\s]([0-9]{3})[\\s]([\u4e00-\u9fa5A-Z]+)[\\s]([0-9\\.]+)[\\s]([0-9\\.])[\\s](((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3})")
	/**
	172.0.0.12 - - [04/Mar/2018:13:49:52 +0000] http "GET /foo?query=t HTTP/1.0" 200 2133 "-" "KeepAliveClient" "-" 1.005 1.854
	*/
	r := regexp.MustCompile(`([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"(.*?)\"\s+\"([\d\.-]+)\"\s+([\d\.-]+)\s+([\d\.-]+)`)

	for v := range l.rc {
		ret := r.FindStringSubmatch(string(v))
		if len(ret) != 14 {
			TypeMonitorChan <- TypeErrNum
			log.Println("FindStringSubmatch error:", string(v))
			continue
		}
		message := &Message{}
		loc, _ := time.LoadLocation("Asia/Shanghai")
		// 2006-01-02 15:04:05
		t, err := time.ParseInLocation("02/Jan/2006:15:04:05 +0000", ret[4], loc)
		if err != nil {
			TypeMonitorChan <- TypeErrNum
			log.Println("ParseInLocation error:", err.Error(), ret[4])
			continue
		}
		// 创建时间
		message.TimeLocal = t

		reqSli := strings.Split(ret[6], " ")
		if len(reqSli) != 3 {
			TypeMonitorChan <- TypeErrNum
			log.Panicln("strings.Split error", ret[6])
			continue
		}
		message.Method = reqSli[0]

		u, err := url.Parse(reqSli[1])
		if err != nil {
			TypeMonitorChan <- TypeErrNum
			log.Panicln("url parse fail:", err)
			continue
		}
		message.Path = u.Path

		message.Scheme = ret[5]
		message.Status = ret[7]

		upsteamTime, _ := strconv.ParseFloat(ret[12], 64)
		requestTime, _ := strconv.ParseFloat(ret[13], 64)
		message.UpstreamTime = upsteamTime
		message.RequestTime = requestTime

		l.wc <- message
	}
}

// echo '172.0.0.12 - - [04/Mar/2018:13:49:52 +0000] http "GET /foo?query=t HTTP/1.0" 200 2133 "-" "KeepAliveClient" "-" 1.005 1.854' >> concurrent/error.log
// 写入日志可以做出自动任务

func main() {
	// l := "[2021-04-19T17:51:06.139128+08:00] http://community.xqd.com/api/v1/child?id=1&name=simba 11068 3f29e074851a96ccb517a1699613e039 {\"name\":\"simba\"} 401 未登录 0.56615 0 172.22.0.1"
	// r := regexp.MustCompile("\\[([\\w-:\\+\\.]*)\\][\\s](htt[ps]://[\\w./\\?=&]+)[\\s]([0-9]+)[\\s]([\\w]+)[\\s]({.+})[\\s]([0-9]{3})[\\s]([\u4e00-\u9fa5A-Z]+)[\\s]([0-9\\.]+)[\\s]([0-9\\.])[\\s](((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3})")
	// s := r.FindStringSubmatch(l)
	// if s == nil {
	// 	fmt.Println("error")
	// 	return
	// }
	// fmt.Println(s)
	// for k, v := range s {
	// 	fmt.Println(k, v)
	// }
	r := &ReadFromFile{
		path: "concurrent/error.log",
	}
	w := &WriteToInfluxDB{
		influxDBDsn: "username",
	}
	lp := &LogProcess{
		rc:    make(chan []byte, 200),
		wc:    make(chan *Message, 200),
		read:  r,
		write: w,
	}

	go lp.read.Read(lp.rc)
	for i := 0; i < 2; i++ {
		go lp.Process()
	}
	for i := 0; i < 5; i++ {
		go lp.write.Write(lp.wc)
	}
	m := &Monitor{
		startTime: time.Now(),
		data:      SystemInfo{},
	}
	m.start(lp)
}
