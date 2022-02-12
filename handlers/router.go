package handlers

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	// load css
	r.Static("/static/css", "./static/css")
	// load javascripts
	r.Static("/static/js", "./static/js")
	// load html
	r.LoadHTMLGlob("static/html/*")

	globalPath := r.Group("/twitter")
	// ping test
	globalPath.GET("/ping", Ping)
	// home page
	globalPath.GET("/index", Index)
	initPageRouters(globalPath)
	initAPIRouters(globalPath)
}

func initPageRouters(gp *gin.RouterGroup) {
	userGroup := gp.Group("user")
	// user login page
	userGroup.GET("login", Login)
	// user register page
	userGroup.GET("register", Register)
}

func initAPIRouters(gp *gin.RouterGroup) {
	apiGroup := gp.Group("api")

	apiUserGroup := apiGroup.Group("user")
	// login post request handle
	apiUserGroup.POST("login", Login)
	// register post request handle
	apiUserGroup.POST("register", Register)

}
