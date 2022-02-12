package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tangwh1206/twitter/common"
	"github.com/tangwh1206/twitter/core"
	"github.com/tangwh1206/twitter/faults"
	"github.com/tangwh1206/twitter/pkg/jsonx"
	"github.com/tangwh1206/twitter/pkg/redis"
	"github.com/tangwh1206/twitter/uuid"
)

func Register(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodPost:
		postRegister(ctx)
	case http.MethodGet:
		getRegister(ctx)
	default:
		ctx.JSON(http.StatusMethodNotAllowed, common.Response(nil, faults.ErrBadRequest))
	}
}

func getRegister(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", gin.H{})
}

func postRegister(ctx *gin.Context) {
	var err error
	var req core.RegisterReq
	err = ctx.ShouldBind(&req)
	if err != nil {
		log.Printf("bind param fail, err=%v\n", err)
		ctx.JSON(http.StatusBadRequest, common.Response(nil, faults.ErrBadRequest))
		return
	}
	username := strings.TrimSpace(req.Username)
	password := strings.TrimSpace(req.Password)
	// check username whether exist
	key := common.GenRedisKeyUsernameToUid(username)
	value, err := redis.GetClient().Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		log.Printf("redis.Get fail, key=%v, err=%v\n", key, err)
		ctx.JSON(http.StatusInternalServerError, common.Response(nil, faults.ErrInternalError))
		return
	}
	if len(value) > 0 {
		ctx.JSON(http.StatusBadRequest, common.Response(nil, faults.ErrUserAlreadyExist))
		return
	}
	// generate a uuid as uid
	uid, err := uuid.NewUUID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.Response(nil, faults.ErrInternalError))
		return
	}
	user := core.User{
		Username: username,
		Password: password,
	}
	userinfo := jsonx.SafeDump(user)
	kvs := map[string]interface{}{}
	key = common.GenRedisKeyUsernameToUid(username)
	kvs[key] = uid

	key = common.GenRedisKeyUidToUser(uid)
	kvs[key] = userinfo

	succ, err := redis.GetClient().MSetNX(ctx, kvs).Result()
	if err != nil {
		log.Printf("redis.MSetNX fail, kvs=%+v, err=%v\n", kvs, err)
		ctx.JSON(http.StatusInternalServerError, common.Response(nil, faults.ErrInternalError))
		return
	}
	// redis key set fail, dut to connurrency register
	if !succ {
		log.Printf("redis.MSetNX fail, kvs=%v, succ=%v", kvs, succ)
		ctx.JSON(http.StatusInternalServerError, common.Response(nil, faults.ErrSystemBusy))
		return
	}
	// todo-register successfully, redirect to home page
	// ctx.JSON(http.StatusOK, common.Response(nil, nil))
	ctx.Redirect(http.StatusFound, "/twitter/index")
}
