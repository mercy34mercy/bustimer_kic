package redisclient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mercy34mercy/bustimer_kic/bustimer/booststrap"
	"github.com/mercy34mercy/bustimer_kic/bustimer/domain/model"
	redis "github.com/redis/go-redis/v9"
	"time"
)

var (
	Client *redis.Client
	ctx    = context.Background()
)

func NewClient() {
	opt, _ := redis.ParseURL(booststrap.RedisDatabaseURL)
	Client = redis.NewClient(opt)
}

func Set(key string, value model.TimeTable) error {
	serialized, _ := json.Marshal(value)
	err := Client.Set(ctx, key, serialized, time.Hour*24).Err()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) (*model.TimeTable, bool) {
	val, err := Client.Get(ctx, key).Bytes()
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
	if val == nil {
		return nil, false
	}
	deserialized := &model.TimeTable{}
	json.Unmarshal(val, deserialized)
	fmt.Printf("key : %s data found", key)
	return deserialized, true
}
