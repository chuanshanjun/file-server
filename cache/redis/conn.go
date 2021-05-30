package redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

var (
	pool          *redis.Pool
	redisHost     = "115.159.34.56"
	redisPassword = ""
)

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     30,
		MaxActive:   30,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			// 1 打开连接
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				log.Fatalf("dial redis err:%v\n", err)
				return nil, err
			}

			// 2 访问认证
			if _, err = c.Do("AUTH", redisPassword); err != nil {
				c.Close()
				log.Fatalf("do redis auth err:%v\n", err)
				return nil, err
			}

			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}

			_, err := c.Do("PING")
			return err
		},
	}
}

func init() {
	pool = newRedisPool()
}

func RedisPool() *redis.Pool {
	return pool
}
