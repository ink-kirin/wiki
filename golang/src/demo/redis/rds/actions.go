package rds

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
)

var rds redis.Conn

func init() {
	//连接地址
	RedisConn := "r-2ze2r7gdb9i4cim8r2pd.redis.rds.aliyuncs.com:6379"
	//db分区
	RedisDbNum := 2
	//密码
	RedisPassword := "ldIla6cg2Bcdo_ggwb20"
	// 链接池
	pool := &redis.Pool{
		MaxIdle:     8,                 // 最大空闲链接数
		MaxActive:   0,                 // 表示和数据库的最大链接数，0 表示没有限制
		IdleTimeout: 100 * time.Second, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码，链接那个ip的redis
			logs.Info(RedisConn)
			c, err := redis.Dial("tcp", RedisConn)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			if RedisPassword != "" {
				if _, authErr := c.Do("AUTH", RedisPassword); authErr != nil {
					return nil, fmt.Errorf("redis auth password error: %s", authErr)
				}
			}
			//选择分区
			c.Do("SELECT", RedisDbNum)
			return c, nil
		},
		//ping
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
	// 从pool取出一个链接
	rds = pool.Get()
}

func Set(key string, val string) (bool, error) {
	_, err := rds.Do("SET", key, val)
	if err != nil {
		logs.Error("set error", err.Error())
		return false, err
	}
	return true, nil
}

func Get(key string) (string, error) {
	val, err := redis.String(rds.Do("GET", key))
	if err != nil {
		logs.Error("get error", err.Error())
		return "", err
	}
	return val, nil
}
