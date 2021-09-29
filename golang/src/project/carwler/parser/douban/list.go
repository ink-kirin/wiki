package douban

import (
	"project/carwler/engine"
	"regexp"
)

const listRe = `<a href="([^"]+)" title="([^"]+)"[\s\S]*?onclick=`

func ParseList(c []byte) engine.ParserResult {
	re := regexp.MustCompile(listRe)
	match := re.FindAllSubmatch(c, -1)
	res := engine.ParserResult{}
	for _, m := range match {
		name := string(m[2])
		res.Items = append(res.Items, name)
		res.Requests = append(res.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParseDetail(c, name)
			},
		})
	}
	return res
}
