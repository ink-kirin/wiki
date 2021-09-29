package tool

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"path"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	TelExp    = `(1[3456789]\d)(\d{4})(\d{4})`
	EmailExp  = `([a-zA-Z0-9_-]+)@([a-zA-Z0-9_]+)(\.[a-zA-Z0-9_]+)+`
	UrlExp    = `<a[\s\S]+?href=["|'](http[\s\S]+?)["|']`
	Card18Exp = `([1-6][1-9]|50)\d{4}(18|19|20)\d{2}((0[1-9])|10|11|12)(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]`
	Card15Exo = `([1-6][1-9]|50)\d{4}\d{2}((0[1-9])|10|11|12)(([0-2][1-9])|10|20|30|31)\d{3}`
	ImgExp    = `<img[\s\S]+?data-original=["|'](http[\s\S]+?)["|']`
)

// 随机字符串的原始值
var defaultLetters = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890")

// 随机字符串
func RandomString(n int, allowedChars ...[]rune) string {
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

// 获取URL内容
func GetUrl(url string) io.Reader {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	return r.Body
}

// 执行正则
func ExtRegular(s string, reg string) (res [][]string) {
	r := regexp.MustCompile(reg)
	res = r.FindAllStringSubmatch(s, -1)
	// fmt.Println(reflect.TypeOf(res)) // 打印变量类型
	return
}

// 获取列表页中详情页的地址
func GetListInfoUrl(u string) (error, []string) {
	r, err := http.Get(u)
	if err != nil {
		return err, nil
	}
	defer r.Body.Close()
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return err, nil
	}
	var urls []string
	doc.Find(".thumb-container-big a").Each(func(i int, s *goquery.Selection) {
		val, exists := s.Attr("href")
		if exists {
			val = strings.TrimSpace(val)
			if strings.Contains(val, "big.php") {
				urls = append(urls, "https://wall.alphacoders.com/"+val)
			}
		}
	})
	return nil, urls
}

// 下载详情页的大图
func GetImagesDownload(u string, urlMap map[string]bool) error {
	if _, ok := urlMap[u]; !ok {
		r, err := http.Get(u)
		if err != nil {
			return err
		}
		defer r.Body.Close()
		doc, err := goquery.NewDocumentFromReader(r.Body)
		if err != nil {
			return err
		}
		doc.Find(".main-content").Each(func(i int, s *goquery.Selection) {
			val, exists := s.Attr("src")
			if exists {
				val = strings.TrimSpace(val)
				fmt.Println(val)
				DownloadFile(val)
				urlMap[u] = true
				fmt.Println("urlmap:", urlMap)

			}
		})
	}
	return nil
}

// 获取URL图片地址中的后缀
func GetUrlExtension(u string) string {
	ext := path.Base(u) // 1.png
	// fmt.Println(path.Ext(u)) // .png
	if strings.Contains(ext, "?") {
		spExt := strings.Split(ext, "?")
		ext = spExt[0]
	}
	return ext
}

// 下载图片到本地
func DownloadFile(u string) string {
	r, err := http.Get(u)
	if err != nil {
		return "文件地址错误:" + err.Error()
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "获取内容失败" + err.Error()
	}
	filename := "./wallpaper/images/" + GetUrlExtension(u)
	fmt.Println(filename)
	err = ioutil.WriteFile(filename, b, 0777)
	if err != nil {
		fmt.Println()
		return "生成文件失败:" + err.Error()
	}
	return ""
}
