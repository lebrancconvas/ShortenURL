package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lebrancconvas/ShortenURL/controllers"
)

func NewRouter(controller *controllers.Controller) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default()) 

	api := router.Group("/api")
	{
		api.GET("/ping", controller.Ping)
		v1 := api.Group("/v1")
		{
			v1.GET("/urls", controller.GetAllURL)
			v1.GET("/urls/:id", controller.GetURLByID)
			short := v1.Group("/short")
			{
				short.POST("/url", controller.CreateShortURL)
				short.GET("/:short_url", controller.GetShortURL)
				short.GET("/url/:id", controller.GetShortURLByID)
			}
		}
	}

	return router
}
