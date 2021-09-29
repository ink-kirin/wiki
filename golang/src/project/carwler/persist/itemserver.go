package persist

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

func ItemServer() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Server: Got Item #%d : %v", itemCount, item)
			itemCount++
		}
	}()
	return out
}

func Save(item interface{}) {
	config := elasticsearch.Config{
		Addresses: []string{"http://127.0.0.1:9200"},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,          // 最大空闲连接数
			ResponseHeaderTimeout: time.Second, // 读取头部时长
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second, // 连接超时时间
				KeepAlive: 30 * time.Second, // 连接保持超时时间
			}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
			},
		},
	}
	es, err := elasticsearch.NewClient(config)
	if err != nil {
		panic(err)
	}
	var b strings.Builder
	b.WriteString(`{"name":"政治","instructor":"学习","content":"lalalala"}`)
	req := esapi.IndexRequest{
		Index:      "douban",
		DocumentID: "",
		Body:       strings.NewReader(b.String()),
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.IsError() {
		log.Printf("[%s] error ", res.Status())
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("body error : %s", err)
		} else {
			log.Printf("[%s] %s, version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
}
