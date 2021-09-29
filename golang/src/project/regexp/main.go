package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"math/rand"
	"net/http"
	"path"
	"regexp"
	"strings"
	"sync"
)

const (
	Tel    = `(1[3456789]\d)(\d{4})(\d{4})`
	Email  = `([a-zA-Z0-9_-]+)@([a-zA-Z0-9_]+)(\.[a-zA-Z0-9_]+)+`
	Url    = `<a[\s\S]+?href=["|'](http[\s\S]+?)["|']`
	Card18 = `([1-6][1-9]|50)\d{4}(18|19|20)\d{2}((0[1-9])|10|11|12)(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]`
	Card15 = `([1-6][1-9]|50)\d{4}\d{2}((0[1-9])|10|11|12)(([0-2][1-9])|10|20|30|31)\d{3}`
	Img    = `<img[\s\S]+?data-original=["|'](http[\s\S]+?)["|']`
)

func getHtml(url string) string {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func crawl(url string, reg string) (res [][]string) {
	s := getHtml(url)
	r := regexp.MustCompile(reg)
	res = r.FindAllStringSubmatch(s, -1)
	// fmt.Println(reflect.TypeOf(res)) // 打印变量类型
	return
}

var defaultLetters = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890")

func randomString(n int, allowedChars ...[]rune) string {
	var letters []rune
	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func downloadImg(u string) {
	r, err := http.Get(u)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	ext := path.Ext(u)
	if strings.Contains(ext, "?") {
		spExt := strings.Split(ext, "?")
		ext = spExt[0]
	}

	filename := "./images/" + randomString(18) + ext
	err = ioutil.WriteFile(filename, b, 0777)
	if err != nil {
		panic(err)
	}
	fmt.Println("下载成功")
}

func main() {
	// tel
	// s := crawl("https://www.xuanxuanhao.com/", Tel)
	// for k, v := range s {
	// 	fmt.Println(k, v)
	// }

	// email

	// url

	// Id Card
	// s2 := crawl("https://m.thepaper.cn/baijiahao_12144816", Card18)

	// img
	s := crawl("https://www.163.com/", Img)
	var wg sync.WaitGroup
	var ch = make(chan int, 3)
	for k, v := range s {
		u := v[1]
		fmt.Println(k, u)
		wg.Add(1)
		go func(k int) {
			ch <- k
			downloadImg(u)
			<-ch
			wg.Done()
		}(k)
	}
	wg.Wait()
}
