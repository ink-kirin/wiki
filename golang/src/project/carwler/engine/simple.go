package engine

import (
	"log"
	"project/carwler/fetcher"
)

// 单列
type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	// 维持一个request队列
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		// 取出队列中的第一个request请求
		r := requests[0]
		requests = requests[1:]
		// 爬取数据
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		// 将爬取结果里的request请求继续加到request队列总
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v \n", item)
		}
	}
}

func worker(r Request) (ParserResult, error) {
	log.Printf("url: %v", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetcher url %s: %v \n", r.Url, err)
		return ParserResult{}, err
	}
	// 解析爬取的结果
	return r.ParserFunc(body), nil
}
