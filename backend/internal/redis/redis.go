package redis

import (
	"fmt"
	"github/Omar-Radwan/backend/internal/constants"
	"os"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

var REDIS_COUNT int64 = 0
var client redis.Conn = nil

func Connect() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 120000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", os.Getenv(constants.REDIS_SERVICE_NAME), os.Getenv(constants.REDIS_SERVICE_PORT)))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func InsertAndCount(increment bool) int64 {
	if client == nil {
		client = Connect().Get()
	}
	if increment {
		client.Do("LPUSH", "mylist", strconv.Itoa(int(REDIS_COUNT)))
		REDIS_COUNT++
	}
	result, _ := client.Do("LLEN", "mylist")
	return result.(int64)
}
