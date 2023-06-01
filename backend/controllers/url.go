package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lebrancconvas/ShortenURL/models"
)

type Controller struct {
	Model *models.Model
}

func NewController(model *models.Model) *Controller {
	return &Controller{Model: model}
}

func (c *Controller) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}


func (c *Controller) CreateShortURL(ctx *gin.Context) {

}

func (c *Controller) GetShortURL(ctx *gin.Context) {

}


func (c *Controller) GetAllURL(ctx *gin.Context) {
	urls, err := c.Model.GetAllURL()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get all url: " + err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, urls)
}
