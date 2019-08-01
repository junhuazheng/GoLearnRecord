package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func ConnectRedis() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn reddis failde, ", err)
		return
	}

	defer c.Close()
}