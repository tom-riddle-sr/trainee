package services

import (
	"encoding/json"
	"fmt"
	"trainee/fibertrainee3/database/redis"
	db0 "trainee/fibertrainee3/model/entity/redis"
	"trainee/fibertrainee3/model/input"
	"trainee/fibertrainee3/repository"
)

type IServicesRedis interface {
	Set(values input.SetRedisData) error
	Get(values input.RedisKeyRequest) error
	Del(values input.RedisKeyRequest) error
}
type ServicesRedis struct {
	repo *repository.Repo
}

func NewServicesRedis(repo *repository.Repo) IServicesRedis {
	return &ServicesRedis{
		repo: repo,
	}
}

func (s *ServicesRedis) Set(values input.SetRedisData) error {
	data := input.SetRedisData{
		Name:  values.Name,
		Phone: values.Phone,
	}
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := s.repo.RedisRepo.Set(redis.GetRedisDB(), "userData", string(json), 0); err != nil {
		return err
	}
	return nil
}
func (s *ServicesRedis) Get(values input.RedisKeyRequest) error {
	model := &db0.RedisStruct{}
	err := s.repo.RedisRepo.Get(redis.GetRedisDB(), values.Key, model)
	if err != nil {
		return err
	}
	return nil
}
func (s *ServicesRedis) Del(values input.RedisKeyRequest) error {
	fmt.Print(values.Key)
	err := s.repo.RedisRepo.Delete(redis.GetRedisDB(), values.Key)
	if err != nil {
		return err
	}
	return nil
}
