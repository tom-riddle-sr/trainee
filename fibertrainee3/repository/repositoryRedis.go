package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type IRedisRepo interface {
	Set(c *redis.Client, key string, value interface{}, expirationTime int32) error
	SetM(c *redis.Client, values map[string]interface{}) error
	Get(c *redis.Client, key string, model interface{}) error
	SAdd(c *redis.Client, key string, values ...interface{}) error
	SIsMember(c *redis.Client, key string, value interface{}) (bool, error)
	SMember(c *redis.Client, key string) ([]string, error)
	LPush(c *redis.Client, key string, values ...interface{}) error
	RPush(c *redis.Client, key string, values ...interface{}) error
	LRange(c *redis.Client, key string, start, stop int64) ([]string, error)
	ZAdd(c *redis.Client, key string, values ...redis.Z) error
	ZRangeWithScores(c *redis.Client, key string, start, stop int64) ([]redis.Z, error)
	Delete(c *redis.Client, key ...string) error
	Test()
}

type RedisRepo struct{}

func NewRedisRepo() IRedisRepo {
	return &RedisRepo{}

}

// String
// SetNX 只有在key不存在時才會設置value
func (r *RedisRepo) Set(c *redis.Client, key string, value interface{}, expirationTime int32) error {
	err := c.Set(key, value, time.Duration(expirationTime)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepo) SetM(c *redis.Client, values map[string]interface{}) error {
	err := c.MSet(values).Err()
	if err != nil {
		return err
	}
	return nil
}

// map[string]interface{}{
// 	"key1": "value1",
// 	"key2": "value2",
// 	"key3": "value3",
// }

func (r *RedisRepo) Get(c *redis.Client, key string, model interface{}) error {
	value, err := c.Get(key).Result()
	if err == redis.Nil {
		fmt.Println(key + " does not exist")
		return err
	} else if err != nil {
		return err
	}
	if err = json.Unmarshal([]byte(value), &model); err != nil {
		return err
	}

	return nil
}

// Set
func (r *RedisRepo) SAdd(c *redis.Client, key string, values ...interface{}) error {
	err := c.SAdd(key, values...).Err()
	if err != nil {
		return err
	}
	return nil
}

// 檢查一個元素是否存在於集合中
func (r *RedisRepo) SIsMember(c *redis.Client, key string, value interface{}) (bool, error) {
	isMember, err := c.SIsMember(key, value).Result()
	if err != nil {
		return false, err
	}
	return isMember, nil
}

// 獲取集合中的所有元素
func (r *RedisRepo) SMember(c *redis.Client, key string) ([]string, error) {
	members, err := c.SMembers(key).Result()
	if err != nil {
		return nil, err
	}
	return members, nil
}

// List
// L前面,R後面
func (r *RedisRepo) LPush(c *redis.Client, key string, values ...interface{}) error {
	err := c.LPush(key, values...).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepo) RPush(c *redis.Client, key string, values ...interface{}) error {
	err := c.RPush(key, values...).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepo) LRange(c *redis.Client, key string, start, stop int64) ([]string, error) {
	values, err := c.LRange(key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	return values, nil
}

// Sorted Set
func (r *RedisRepo) ZAdd(c *redis.Client, key string, values ...redis.Z) error {
	err := c.ZAdd(key, values...).Err()
	if err != nil {
		return err
	}
	return nil
}

// 獲取有序集合的所有元素，按分數排序
func (r *RedisRepo) ZRangeWithScores(c *redis.Client, key string, start, stop int64) ([]redis.Z, error) {
	values, err := c.ZRangeWithScores(key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	return values, nil
}

func (r *RedisRepo) Delete(c *redis.Client, key ...string) error {
	fmt.Println(key)
	err := c.Del(key...).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *RedisRepo) Test() {
	fmt.Println("Redis Test")
}
