package redis

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/tangwh1206/twitter/conf"
	"github.com/tangwh1206/twitter/core"
)

var (
	Nil = redis.Nil
)

var (
	client *redis.Client
	once   sync.Once
)

func Init(config *core.RedisConfig) {
	once.Do(
		func() {
			addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
			options := &redis.Options{
				Addr:     addr,
				Password: "", // no password set
				DB:       0,  // use default db
			}
			client = redis.NewClient(options)
		},
	)
}

func GetClient() *redis.Client {
	if client == nil {
		Init(&conf.GetSetting().Redis)
	}
	return client
}
