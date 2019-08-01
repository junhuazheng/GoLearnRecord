package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func Expire() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed, ", err)
		return
	}
	defer c.Close()

	_, err = c.Do("expire", "abc", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
}