package session

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/tangwh1206/twitter/common"
	"github.com/tangwh1206/twitter/pkg/redis"
	"github.com/tangwh1206/twitter/pkg/utils"
)

const (
	sessionSaltPrefix = "QUKQSWJH"
	sessionSaltSuffix = "FOWIFJDS"
)

var (
	ErrNil = errors.New("session is nil")
)

func GenSessionByUID(uid int64) string {
	// todo-dynamic gen seission
	if uid <= 0 {
		return ""
	}
	uidStr := strconv.FormatInt(uid, 10)
	text := strings.Join([]string{sessionSaltPrefix, uidStr, sessionSaltSuffix}, "_")
	session := utils.GenMD5(text)
	return session
}

func GetUIDBySession(ctx context.Context, session string) (int64, error) {
	key := common.GenRedisKeySessionId2Uid(session)
	value, err := redis.GetClient().Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, ErrNil
		} else {
			return 0, err
		}
	}
	uid, err := strconv.Atoi(value)
	if err != nil {
		return 0, ErrNil
	}
	return int64(uid), nil
}
