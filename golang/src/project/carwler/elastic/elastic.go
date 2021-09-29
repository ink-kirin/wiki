package elastic

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"net"
	"net/http"
	"time"
)

type ElasticParam struct {
	IndexName string
	Doc       map[string]interface{}
	ID        string
}

var ES *elasticsearch.Client

func init() {
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
	client, err := elasticsearch.NewClient(config)
	if err != nil {
		panic(err)
	}
	ES = client
}

// 创建或更文档
func (e ElasticParam) Index() (bool, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(e.Doc); err != nil {
		return false, err
	}
	fmt.Println(e.IndexName)
	res, err := ES.Index(e.IndexName, &buf)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	body, err := verifyBody(res, "created")
	return body, err
}

// 搜索
func (e ElasticParam) Search() {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(e.Doc); err != nil {
		panic(err)
	}
	res, err := ES.Search(ES.Search.WithIndex(e.IndexName),
		ES.Search.WithContext(context.Background()),
		ES.Search.WithBody(&buf),
		ES.Search.WithTrackTotalHits(true),
		ES.Search.WithPretty())
	if err != nil {
		panic(nil)
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// 添加文档（需要知道_id, _id存在返回409）
func (e ElasticParam) Create() (bool, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(e.Doc); err != nil {
		return false, err
	}
	res, err := ES.Create(e.IndexName, e.ID, &buf)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	body, err := verifyBody(res, "created")
	return body, err
}

// 通过ID获取文档
func Get() {

}

// 通过ID更新文档
func (e ElasticParam) Update() (bool, error) {
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"doc": e.Doc,
	}
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		return false, err
	}
	res, err := ES.Update(e.IndexName, e.ID, &buf)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	body, err := verifyBody(res, "updated")
	return body, err
}

// 通过匹配条件更新文档
func (e ElasticParam) UpdateByQuery() {
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": "不动明王",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		panic(err)
	}
	res, err := ES.UpdateByQuery([]string{e.IndexName},
		ES.UpdateByQuery.WithBody(&buf),
		ES.UpdateByQuery.WithContext(context.Background()),
		ES.UpdateByQuery.WithPretty())
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// 通过ID删除文档
func (e ElasticParam) Delete() (bool, error) {
	res, err := ES.Delete(e.IndexName, e.ID)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	body, err := verifyBody(res, "deleted")
	return body, err
}

// 通过匹配条件删除文档
func (e ElasticParam) DeleteByQuery() {

}

func verifyBody(res *esapi.Response, cmd string) (bool, error) {
	log.Printf("Delete Res Body Data:%v", res.String())
	if res.IsError() {
		log.Printf("[%s] Error", res.Status())
		return false, errors.New("Error " + res.Status())
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
			return false, err
		} else {
			successful := int(r["_shards"].(map[string]interface{})["successful"].(float64))
			if r["result"] == cmd && successful == 1 {
				return true, nil
			}
			return false, nil
		}
	}
}
