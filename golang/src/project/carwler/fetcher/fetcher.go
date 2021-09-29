package fetcher

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var miterlimit = time.Tick(2000 * time.Millisecond)

/**
请求URL网络资源
*/
func Fetch(u string) ([]byte, error) {
	<-miterlimit
	tr := &http.Transport{TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	}}

	proxyUrl, err := url.Parse("http://127.0.0.1:7890")
	if err == nil { // 使用传入代理
		tr.Proxy = http.ProxyURL(proxyUrl)
	}

	//生成client 参数为默认
	client := &http.Client{Transport: tr}
	//提交请求
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		fmt.Printf("NewRequest Error:%v:\n", err)
		return nil, err
	}
	//增加header选项
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (HTML, like Gecko) Chrome/90.0.4430.85 Safari/537.36")
	//处理返回结果
	r, err := client.Do(req)
	if err != nil {
		fmt.Printf("Client Do Error:%v:\n", err)
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		fmt.Printf("Error status code Error:%d\n", r.StatusCode)
		return nil, err
	}
	//编码转换，自动检测网页编码
	bodyReader := bufio.NewReader(r.Body)
	e := DetermineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	b, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		fmt.Printf("ReadAll Error:%d\n", r.StatusCode)
		return nil, err
	}
	fmt.Printf("fetch %s\n", u)
	return b, nil
}

/*
检测html页面编码
*/
func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	//这里的r读取完得保证resp.Body还可读
	bytes, err := r.Peek(1024)
	//如果解析编码类型时遇到错误,返回UTF-8
	if err != nil {
		log.Printf("Fetcher err:%v\n", err)
		return unicode.UTF8
	}
	//这里简化,不取是否确认
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
