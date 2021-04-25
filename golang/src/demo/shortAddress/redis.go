package main

import (
	"time"

	"github.com/go-redis/redis/v8"
)

// Redis Key
const (
	// global 全局计数器
	URLIDKEY = "next.url.id"
	// 短地址与长地址之间的关系
	ShortlinkKey = "shortlink:%s:url"
	// 长地址的hash与短地址之间的关系
	URLHashKey = "urlhash:%s:url"
	// 短地址详细信息
	ShortlinkDetailKey = "shortlink:%s:detail"
)

// RedisCli

type RedisCli struct {
	Cli *redis.Client
}

// URLDetail
type URLDetail struct {
	URL                 string        `json:"url"`
	CreatedAt           string        `json:"created_at"`
	ExpirationInMinutes time.Duration `json:"expiration_in_minutes"`
}

// 初始化Redis结构体
func NewRedisCli(addr string, passwd string, db int) *RedisCli {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})
	if _, err := c.Ping(c.Context()).Result(); err != nil {
		panic(err)
	}
	return &RedisCli{Cli: c}
}
