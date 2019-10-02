package api

import (
	"net/http"

	"github.com/FishwithCat/gz-backend/pkg/util"
	"github.com/FishwithCat/gz-backend/src/models"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	Username string
	Password string
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// auth := Auth{Username: username, Password: password}
	isExistUser := models.CheckAuth(username, password)
	if !isExistUser {
		c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{
			"msg": "Auth failed",
		})
		return
	}
	token, err := util.GenerateToken(username, password)
	if err != nil {
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"token": token,
	})
}
