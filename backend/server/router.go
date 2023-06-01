package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lebrancconvas/ShortenURL/controllers"
)

func NewRouter(controller *controllers.Controller) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/ping", controller.Ping)
		v1 := api.Group("/v1")
		{
			v1.GET("/urls", controller.GetAllURL)
			short := v1.Group("/short")
			{
				short.POST("/url", controller.CreateShortURL)
				short.GET("/url", controller.GetShortURL)
			}
		}
	}

	return router
}
