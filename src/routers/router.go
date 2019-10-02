package routers

import (
	"github.com/FishwithCat/gz-backend/src/api"
	jwt "github.com/FishwithCat/gz-backend/src/middware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.GET("/auth", api.GetAuth)

	apiv1 := router.Group("/api/v1/")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	return router
}
