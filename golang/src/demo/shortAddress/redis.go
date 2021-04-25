package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pilu/go-base62"
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
	if _, err := c.Ping().Result(); err != nil {
		panic(err)
	}
	return &RedisCli{Cli: c}
}

func (r *RedisCli) Shorten(url string, exp int64) (string, error) {
	h := toSha1(url)

	d, err := r.Cli.Get(fmt.Sprintf(URLHashKey, h)).Result()
	if err == redis.Nil {
		// 不存在不做处理
	} else if err != nil {
		return "", err
	} else {
		if d == "{}" {
			// 为空不做处理
		} else {
			return d, nil
		}
	}
	// 第一次转换存入redis
	err = r.Cli.Incr(URLIDKEY).Err()
	if err != nil {
		return "", err
	}

	id, err := r.Cli.Get(URLIDKEY).Int64()
	if err != nil {
		return "", err
	}

	eid := base62.Encode(int(id))

	err = r.Cli.Set(fmt.Sprintf(ShortlinkKey, eid), url, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}

	err = r.Cli.Set(fmt.Sprintf(URLHashKey, h), eid, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}
	detail, err := json.Marshal(
		&URLDetail{
			URL: url,
			CreatedAt: time.Now().String(),
			ExpirationInMinutes: time.Duration(exp)}
		)
	if err != nil {
		return "", err
	}

	err = r.Cli.Set(fmt.Sprintf(ShortlinkDetailKey, eid), detail, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}
	return eid, nil
}

func toSha1(d string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}