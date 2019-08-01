package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func BatchSet() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed, ", err)
		return
	}
	defer c.Close()

	_, err = c.Do("MSet", "abc", "100", "def", 300)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Ints(c.Do("MGet", "abc", "def"))
	if err != nil {
		fmt.Println("get abc failed, ", err)
		return
	}

	for _, v := range r {
		fmt.Println(v)
	}
}