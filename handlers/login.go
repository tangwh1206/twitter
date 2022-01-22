package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tangwh1206/twitter/common"
	"github.com/tangwh1206/twitter/conf"
	"github.com/tangwh1206/twitter/core"
	"github.com/tangwh1206/twitter/faults"
	"github.com/tangwh1206/twitter/pkg/redis"
	"github.com/tangwh1206/twitter/session"
)

func Login(ctx *gin.Context) {
	var err error
	var req core.LoginReq
	err = ctx.ShouldBind(&req)
	if err != nil {
		log.Printf("bind param fail, err=%v\n", err)
		ctx.JSON(http.StatusBadRequest, common.Response(nil, faults.ErrBadRequest))
		return
	}
	username := strings.TrimSpace(req.Username)
	password := strings.TrimSpace(req.Password)
	key := common.GenRedisKeyUsernameToUid(username)
	value, err := redis.GetClient().Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			// user not register
			log.Printf("user not register, key=%v, err=%v\n", key, err)
			ctx.JSON(http.StatusNonAuthoritativeInfo, common.Response(nil, faults.ErrUserNotFound))
			return
		} else {
			log.Printf("redis.Get fail, key=%v, err=%v\n", key, err)
			ctx.JSON(http.StatusInternalServerError, common.Response(nil, faults.ErrInternalError))
			return
		}
	}
	if len(value) == 0 {
		log.Printf("redis.Get value is empty, key=%v, value=%v\n", key, value)
		ctx.JSON(http.StatusInternalServerError, common.Response(nil, faults.ErrInternalError))
		return
	}
	uid, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("strconv.Atoi fail, value=%v, err=%v\n", value, err)
		ctx.JSON(http.StatusInternalServerError, common.Response(nil, faults.ErrInternalError))
		return
	}

	key = common.GenRedisKeyUidToUser(int64(uid))
	value, err = redis.GetClient().Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			// user info not found
			log.Printf("user info not found, key=%v, err=%v\n", key, err)
			ctx.JSON(http.StatusInternalServerError, common.Response(nil, faults.ErrInternalError))
			return
		} else {
			log.Printf("redis.Get fail, key=%v, err=%v\n", key, err)
			ctx.JSON(http.StatusInternalServerError, common.Response(nil, faults.ErrInternalError))
			return
		}
	}
	if len(value) == 0 {
		log.Printf("redis.Get value is empty, key=%v, value=%v\n", key, value)
		ctx.JSON(http.StatusInternalServerError, common.Response(nil, faults.ErrInternalError))
		return
	}
	var user core.User
	if len(value) > 0 {
		err = json.Unmarshal([]byte(value), &user)
		if err != nil {
			log.Printf("json.Unmarshal fail, key=%v, val=%v, err=%v\n", key, value, err)
			ctx.JSON(http.StatusInternalServerError, common.Response(nil, faults.ErrInternalError))
			return
		}
	}

	if user.Username != username || user.Password != password {
		ctx.JSON(http.StatusBadRequest, common.Response(nil, faults.ErrPasswordError))
		return
	}
	// set cookie
	sessionKey := session.GenSessionByUID(int64(uid))
	ctx.SetCookie(common.TwitterSessionV1Key, sessionKey, common.SessionExpireTime, "/", conf.GetSetting().Server.Domain, false, true)

	// todo-login successfully, redirect to home page and set cookie
	ctx.JSON(http.StatusOK, common.Response(nil, nil))
}
