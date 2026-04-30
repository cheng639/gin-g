package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisConfig struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"`
	Database int    `json:"database" yaml:"database"`
}

var sClient StandaloneClient

type StandaloneClient struct {
	Client redis.Client
	ctx    context.Context
}

func RedisClient() *StandaloneClient {
	return &sClient
}

func (r *StandaloneClient) IsConnection() bool {
	var err error

	_, err = r.Client.Ping(r.ctx).Result()
	if err != nil {
		return false
	}

	return true
}

func (r *StandaloneClient) Close() error {
	return r.Client.Close()
}

func (r *StandaloneClient) Get(key string) ([]byte, error) {
	return r.Client.Get(r.ctx, key).Bytes()
}

func (r *StandaloneClient) RPush(key string, value string) (int64, error) {
	return r.Client.RPush(r.ctx, key, value).Result()
}

func (r *StandaloneClient) LPop(key string) ([]byte, error) {
	return r.Client.LPop(r.ctx, key).Bytes()
}

func (r *StandaloneClient) Incr(key string) (int64, error) {
	return r.Client.Incr(r.ctx, key).Result()
}

func (r *StandaloneClient) SetNX(key string, val interface{}, expiration time.Duration) (bool, error) {
	return r.Client.SetNX(r.ctx, key, val, expiration).Result()
}

func (r *StandaloneClient) Set(key string, val interface{}, expiration time.Duration) (string, error) {
	return r.Client.Set(r.ctx, key, val, expiration).Result()
}

func (r *StandaloneClient) Del(key string) (int64, error) {
	return r.Client.Del(r.ctx, key).Result()
}

func (r *StandaloneClient) Exists(key string) (int64, error) {
	return r.Client.Exists(r.ctx, key).Result()
}

func (r *StandaloneClient) Do(args ...interface{}) (interface{}, error) {
	return r.Client.Do(r.ctx, args...).Result()
}

func (r *StandaloneClient) HGet(key, field string) (string, error) {
	return r.Client.HGet(r.ctx, key, field).Result()
}

func (r *StandaloneClient) HSet(key, field string, value interface{}) (int64, error) {
	return r.Client.HSet(r.ctx, key, field, value).Result()
}

func (r *StandaloneClient) HExists(key, field string) (bool, error) {
	return r.Client.HExists(r.ctx, key, field).Result()
}
func (r *StandaloneClient) Expire(key string, expiration time.Duration) (bool, error) {
	return r.Client.Expire(r.ctx, key, expiration).Result()
}

// redis zset
func (r *StandaloneClient) ZAddArgs(key string, args redis.ZAddArgs) (int64, error) {
	return r.Client.ZAddArgs(r.ctx, key, args).Result()
}

func (r *StandaloneClient) ZCount(key, min, max string) (int64, error) {
	return r.Client.ZCount(r.ctx, key, min, max).Result()
}
