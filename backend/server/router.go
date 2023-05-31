package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lebrancconvas/ShortenURL/controllers"
)

func NewRouter(controller *controllers.Controller) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/url", controller.CreateShortURL)
			v1.GET("/url", controller.GetShortURL)
		}
	}

	return router
}
