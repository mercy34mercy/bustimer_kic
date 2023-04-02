package redisclient

import (
	"context"
	redis "github.com/redis/go-redis/v9"
)

var (
	Client *redis.Client
	ctx    = context.Background()
)

func NewClient() {
	opt, _ := redis.ParseURL("redis://default:********@apn1-fun-slug-35150.upstash.io:35150")
	Client = redis.NewClient(opt)
}

func Set(key, value string) error {
	err := Client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) (string, error) {
	val, err := Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
