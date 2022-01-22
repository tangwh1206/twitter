package handlers

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	globalPath := r.Group("/twitter")
	globalPath.GET("/ping", Ping)
	globalPath.POST("/login", Login)
	globalPath.POST("/register", Register)

	apiGroup := globalPath.Group("api", nil, nil)
	twitterRouters(apiGroup)
}

func twitterRouters(api *gin.RouterGroup) {
	reader := api.Group("/", nil)
	reader.GET("/post/detail")
}
