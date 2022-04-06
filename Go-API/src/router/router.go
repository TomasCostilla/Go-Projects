package router

import (
	"api.com/src/services"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.GET("/users", services.GetUsers)

	router.Run("localhost:8080")
}
