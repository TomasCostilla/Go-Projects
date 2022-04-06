package models

import (
	"api.com/src/services"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/")
}

func GetUserRoute(router *gin.Engine) {
	router.GET("/", services.GetUsers)
}
