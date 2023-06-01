package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// "github.com/lebrancconvas/ShortenURL/forms"
	"github.com/lebrancconvas/ShortenURL/models"
	"github.com/lebrancconvas/ShortenURL/utils"
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
	type RequestData struct {
		OriginalURL string `json:"original_url" binding:"required"`
	}

	var req RequestData

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	exist, err := c.Model.CheckExistOriginalURL(req.OriginalURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to check url: " + err.Error(),
		})
		return
	}

	if !exist {
		err = c.Model.StoreURL(req.OriginalURL)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to store url: " + err.Error(),
			})
			return
		}

		shortURL := utils.GenerateShortURL()
		err = c.Model.UpdateNewShortenURL(req.OriginalURL, shortURL)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to update short url: " + err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"original_url": req.OriginalURL,
			"short_url": shortURL,
			"message": "Successfully created new short url",
		})
	} else {
		shortURL, err := c.Model.GetShortURLByOriginalURL(req.OriginalURL)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to get short url: " + err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "URL already exist",
			"original_url": req.OriginalURL,
			"short_url": shortURL,
		})
	}
}

func (c *Controller) GetShortURL(ctx *gin.Context) {
	shortURL := ctx.Param("short_url")

	url, err := c.Model.GetOriginalURLByShortURL(shortURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get url: " + err.Error(),
		})
	}

	ctx.Redirect(http.StatusMovedPermanently, url)
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
