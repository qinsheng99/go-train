package redisClient

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

type Redis struct {
	RE *redis.Client
}

type RE struct {
	Name     string `json:"name" form:"name"`
	Data     string `json:"data" form:"data"`
	Stop     int64  `json:"stop" form:"stop"`
	Start    int64  `json:"start" form:"start"`
	PushType string `json:"pushType" form:"pushType"`
}

const HSETKEYS = "hset_%s"

func GetKey(key string, args ...interface{}) string {
	return fmt.Sprintf(key, args...)
}

func GetRedis() (*Redis, error) {
	conn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})
	s, err := conn.Ping().Result()
	if err != nil {
		fmt.Print(s)
		return nil, err
	}
	return &Redis{
		RE: conn,
	}, nil
}
