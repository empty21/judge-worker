package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"judger/pkg/config"
	"time"
)

var client *redis.Client
var ctx = context.TODO()

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisUri,
		Password: config.Config.RedisPassword,
		DB:       config.Config.RedisDatabase,
	})
}

func Get(key string) interface{} {
	value, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}
	return value
}

func Set(key string, value interface{}, ttl int) error {
	var timestamp int64 = 0
	if ttl > 0 {
		timestamp = time.Now().UnixMilli() + int64(1000*ttl)
	}
	return client.Set(ctx, key, value, time.Duration(timestamp)).Err()
}

func KeyExisted(key string) bool {
	value := Get(key)
	return value != nil
}

func WithMachineID(key string) string {
	return config.Config.MachineId + "/" + key
}
