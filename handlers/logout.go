package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Logout(ctx *gin.Context) {
	log.Println("Login called")
}
