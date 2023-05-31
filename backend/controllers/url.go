package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lebrancconvas/ShortenURL/models"
)

type Controller struct {
	Model *models.Model
}

func NewController(model *models.Model) *Controller {
	return &Controller{Model: model}
}


func (c *Controller) CreateShortURL(ctx *gin.Context) {

}

func (c *Controller) GetShortURL(ctx *gin.Context) {
	
}
