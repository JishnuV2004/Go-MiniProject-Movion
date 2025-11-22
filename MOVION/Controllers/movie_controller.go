package controllers

import (
	config "movion/Config"
	models "movion/Models"
	services "movion/Services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateMovie(c *gin.Context){
	var inputMovie models.Movie
	if err := c.ShouldBindJSON(&inputMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid input",
		})
		return
	}
	movie, err := services.CreateMovie(&inputMovie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"movie creation failed",
			"err":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"movie":movie,
	})
}

func GetAllMoviesPage(c *gin.Context) {
    var movies []models.Movie
    config.DB.Find(&movies)
    c.HTML(200, "home.html", gin.H{
        "Movies": movies,
    })
}


func GetMovie(c *gin.Context){
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid request",
		})
		return
	}
	movie, err := services.GetMovie(uint(id64))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"movie":movie,
	})
}
func EditMovie(c *gin.Context){
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid request",
		})
		return
	}

	var input models.Movie
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invaled request",
		})
		return
	}
	movie, err := services.EditMovie(uint(id64), &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"movie":movie,
	})
}

func DeleteMovie(c *gin.Context){
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid request",
		})
		return
	}
	if err := services.DeleteMovie(uint(id64)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"failed movie deleting",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"delete successful",
	})
}