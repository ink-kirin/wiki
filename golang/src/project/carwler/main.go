package main

import (
	"project/carwler/elastic"
)

func main() {

	doc := map[string]interface{}{
		"name": "下次见面",
	}
	e := elastic.ElasticParam{
		IndexName: "douban",
		Doc:       doc,
		ID:        "1",
	}
	e.Search()
	return

	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.QueuedScheduler{},
	//	WorkerCount: 10,
	//	ItemChan:    persist.ItemServer(),
	//}
	//e.Run(engine.Request{
	//	Url:        "https://book.douban.com/",
	//	ParserFunc: douban.ParseTag,
	//})
}
