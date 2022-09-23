package data

import (
	"context"
	"github.com/go-redis/redis/v8"
	c "go_blog/config"
)

// redis 配置读取

var Rdb *redis.Client

func initRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     c.Config.Redis.Host + ":" + c.Config.Redis.Port,
		Password: c.Config.Redis.Password,
		DB:       c.Config.Redis.Database,
	})
	var ctx = context.Background()
	_, err := Rdb.Ping(ctx).Result()

	if err != nil {
		panic("redis connection failed" + err.Error())
	}
}
