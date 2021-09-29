package main

import (
	"project/Carwlers/engine"
	"project/Carwlers/parser/douban"
	"project/Carwlers/persist"
	"project/Carwlers/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemServer(),
	}
	e.Run(engine.Request{
		Url:        "https://book.douban.com/",
		ParserFunc: douban.ParseTag,
	})
}
