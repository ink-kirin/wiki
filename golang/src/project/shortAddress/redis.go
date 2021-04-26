package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pilu/go-base62"
)

// Redis Key
const (
	// global 全局自增器
	URLIDKEY = "next.url.id"
	// 短地址与地址的映射
	ShortlinkKey = "shortlink:%s:url"
	// 地址hash与短地址的映射
	URLHashKey = "urlhash:%s:url"
	// 短地址与详情的映射
	ShortlinkDetailKey = "shortlink:%s:detail"
)

// RedisCli

// top-level context for incoming requests. 请求的顶级上下文. 不知道有啥用,但必须填写
var ctx = context.Background()

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
	if _, err := c.Ping(context.TODO()).Result(); err != nil {
		panic(err)
	}
	return &RedisCli{Cli: c}
}

// Shorten convert url to shortlink
func (r *RedisCli) Shorten(url string, exp int64) (string, error) {
	// convert url to sha1 hash
	h := toSha1(url)

	// fetch it if the url is cached
	d, err := r.Cli.Get(ctx, fmt.Sprintf(URLHashKey, h)).Result()
	if err == redis.Nil {
		// not existed, nothing to do
	} else if err != nil {
		return "", err
	} else {
		if d == "{}" {
			// expiration, nothing to do
		} else {
			return d, nil
		}
	}
	// increase the glabal counter
	err = r.Cli.Incr(ctx, URLIDKEY).Err()
	if err != nil {
		return "", err
	}

	// encode global counter to base62
	id, err := r.Cli.Get(ctx, URLIDKEY).Int64()
	if err != nil {
		return "", err
	}

	eid := base62.Encode(int(id))

	// store the url against this ancoded id
	err = r.Cli.Set(ctx, fmt.Sprintf(ShortlinkKey, eid), url, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}

	// store the url against the hash of it
	err = r.Cli.Set(ctx, fmt.Sprintf(URLHashKey, h), eid, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}
	detail, err := json.Marshal(
		&URLDetail{
			URL:                 url,
			CreatedAt:           time.Now().String(),
			ExpirationInMinutes: time.Duration(exp),
		},
	)
	if err != nil {
		return "", err
	}

	// store the url detail against this encoded id
	err = r.Cli.Set(ctx, fmt.Sprintf(ShortlinkDetailKey, eid), detail, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}
	return eid, nil
}

// ShortlinkInfo returns the detail of the shortlink
func (r *RedisCli) ShortlinkInfo(eid string) (interface{}, error) {
	d, err := r.Cli.Get(ctx, fmt.Sprintf(ShortlinkDetailKey, eid)).Result()
	if err == redis.Nil {
		return "", StatusError{404, errors.New("unknown short URL")}
	} else if err != nil {
		return "", err
	} else {
		return d, nil
	}
}

// Unshorten convert shrotlink to url
func (r *RedisCli) Unshorten(eid string) (string, error) {
	url, err := r.Cli.Get(ctx, fmt.Sprintf(ShortlinkKey, eid)).Result()
	if err == redis.Nil {
		return "", StatusError{404, err}
	} else if err != nil {
		return "", err
	} else {
		return url, nil
	}
}

func toSha1(d string) string {
	h := sha1.New()
	h.Write([]byte(d))
	return hex.EncodeToString(h.Sum(nil))
}
