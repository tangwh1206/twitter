package uuid

import (
	"context"

	"github.com/tangwh1206/twitter/pkg/redis"
)

const (
	RedisKeyUUID = "uuid"
)

func NewUUID(ctx context.Context) (int64, error) {
	return redis.GetClient().Incr(ctx, RedisKeyUUID).Result()
}
