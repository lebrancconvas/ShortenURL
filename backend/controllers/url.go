package controllers

import (
	"log"
	"net/http"
	"strconv"

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

func (c *Controller) GetShortURLByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	shortURL, err := c.Model.GetShortURLByID(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get short url: " + err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, shortURL)
}

func (c *Controller) GetURLByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	url, err := c.Model.GetURLDetailByID(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get url: " + err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, url)
}


func (c *Controller) GetAllURL(ctx *gin.Context) {
	urls, err := c.Model.GetAllURLDetail()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get all url: " + err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, urls)
}
