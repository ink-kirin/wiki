package douban

import (
	"project/Carwlers/engine"
	"project/Carwlers/model"
	"regexp"
	"strconv"
)

var (
	authorReg = regexp.MustCompile(`<span class="pl"> 作者</span>:[\s\S]+?<a.*?>([^<]+)</a>`)
	publicReg = regexp.MustCompile(`<span class="pl">出版社:</span> ([^<]+)<br/>`)
	pageReg   = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
	priceReg  = regexp.MustCompile(`<span class="pl">定价:</span> ([^<]+)<br/>`)
	scoreReg  = regexp.MustCompile(`<strong class="ll rating_num " property="v:average"> ([^<]+)</strong>`)
	intoReg   = regexp.MustCompile(`<div class="intro">[\s\S]+?<p>([^<]+)</p></div>`)
)

func ParseDetail(c []byte, name string) engine.ParserResult {
	detail := model.DoubanDetail{}
	detail.Title = name
	detail.Author = ExtraString(c, authorReg)
	detail.Publicer = ExtraString(c, publicReg)
	page, err := strconv.Atoi(ExtraString(c, pageReg))
	if err == nil {
		detail.Bookpages = page
	}
	detail.Price = ExtraString(c, priceReg)
	detail.Score = ExtraString(c, scoreReg)
	detail.Into = ExtraString(c, intoReg)
	result := engine.ParserResult{
		Items: []interface{}{detail},
	}
	return result
}
func ExtraString(content []byte, reg *regexp.Regexp) string {
	s := reg.FindStringSubmatch(string(content))
	if len(s) >= 2 {
		return s[1]
	} else {
		return ""
	}
}
