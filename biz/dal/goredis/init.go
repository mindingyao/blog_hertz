package goredis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisConn *redis.Pool

func Init() {
	RedisConn = &redis.Pool{
		MaxIdle:     20,
		MaxActive:   30,
		IdleTimeout: 200,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
