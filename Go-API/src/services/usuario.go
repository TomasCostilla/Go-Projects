package services

import (
	"net/http"

	"api.com/src/models"
	"github.com/gin-gonic/gin"
)

/* obtiene todos los usuarios */
func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Users)
}
