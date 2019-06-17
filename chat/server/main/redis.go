package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool 

func initRedisPool(maxIdle, maxActive initRedisPool, idleTimeout time.Duration, host string) {
	pool = &redis.Pool{
		//initialize the number of connections
		MaxIdle: maxIdle,
		MaxActive: maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", host)
		},
	}
}