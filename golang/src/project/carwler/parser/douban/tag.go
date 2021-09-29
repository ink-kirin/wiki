package douban

import (
	"project/Carwlers/engine"
	"regexp"
)

const tagRe = `<a href="([^"]+)" class="tag">([^<]+)</a>`

func ParseTag(c []byte) engine.ParserResult {
	re := regexp.MustCompile(tagRe)
	match := re.FindAllSubmatch(c, -1)
	res := engine.ParserResult{}
	for _, m := range match {
		res.Items = append(res.Items, string(m[2]))
		res.Requests = append(res.Requests, engine.Request{
			Url:        "https://book.douban.com" + string(m[1]),
			ParserFunc: ParseList,
		})
	}
	return res
}
