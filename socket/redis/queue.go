package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func Queue() {
	c, err　:= redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed, ", err)
		return
	}
	defer c.Close()

	_, err = c.Do("lpush", "book list", "abc", "ceg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.String(c.Do("lpop", "book list"))
	if err != nil {
		fmt.Println("get abc failed, ", err)
		return
	}

	fmt.Println(r)
} 