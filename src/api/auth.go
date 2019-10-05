package api

import (
	"fmt"
	"net/http"

	"github.com/FishwithCat/gz-backend/pkg/util"
	"github.com/FishwithCat/gz-backend/src/models"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	Username string
	Password string
}

func AddAuth(c *gin.Context) {
	var newAuth Auth
	c.BindJSON(&newAuth)
	fmt.Println(newAuth)
	err := models.AddAuth(newAuth.Username, newAuth.Password)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"errorMsg": "register failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// auth := Auth{Username: username, Password: password}
	isExist, err := models.CheckAuth(username, password)
	if !isExist {
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
