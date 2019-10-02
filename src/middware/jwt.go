package jwt

import (
	"net/http"

	"github.com/FishwithCat/gz-backend/pkg/util"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := http.StatusOK
		token := c.Query("token")
		// claims, err := util.ParseToken(token)
		if token == "" {
			code = http.StatusUnauthorized
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				code = http.StatusUnauthorized
			}
		}
		if code != http.StatusOK {
			c.JSON(code, gin.H{
				"msg": "need auth",
			})

			c.Abort()
			return
		}
		c.Next()
	}
}
