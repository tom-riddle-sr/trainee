package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

var _redisDB *redis.Client

func New() error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_redisDB = client

	if _, err := client.Ping().Result(); err != nil {
		fmt.Println("連接到redis 失敗:", err)
		return err
	}
	fmt.Println("連接到redis 成功")
	return nil

}

func GetRedisDB() *redis.Client {
	return _redisDB
}
