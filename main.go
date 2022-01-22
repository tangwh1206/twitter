package main

import (
	"log"

	"github.com/tangwh1206/twitter/conf"
	"github.com/tangwh1206/twitter/pkg/redis"
	"github.com/tangwh1206/twitter/server"
)

func Init() {
	conf.Init()

	redis.Init(&conf.GetSetting().Redis)
}

func main() {
	Init()

	svr := server.New(&conf.GetSetting().Server)
	err := svr.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
