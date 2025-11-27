package controllers

import (
	models "movion/Models"
	services "movion/Services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateShow(c *gin.Context) {
	var input struct {
		MovieID  uint    `json:"movie_id"`
		ScreenID uint    `json:"screen_id"`
		ShowTime string  `json:"show_time"` // expect RFC3339 or custom
		Price    float64 `json:"price"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid request",
		})
		return
	}
	st, err := time.Parse(time.RFC3339, input.ShowTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid showtime format, use RFC3339",
		})
		return
	}
	show := &models.Show{
		MovieID: input.MovieID,
		ScreenID: input.ScreenID,
		ShowTime: st,
		Price: input.Price,
	}
	showCreated, err := services.CreteShow(show)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"show": showCreated,
	})
}

func GetShowByID(c *gin.Context) {
idStr := c.Param("id")
id64, err := strconv.ParseUint(idStr, 10, 64)
if err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
return
}
show, err := services.GetShowByID(uint(id64))
if err != nil {
c.JSON(http.StatusNotFound, gin.H{"error": "show not found"})
return
}
c.JSON(http.StatusOK, gin.H{"show": show})
}

func GetAllShows(c *gin.Context){
	shows, err := services.GetAllShows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"shows not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shows":shows,
	})
}

func EditShow(c *gin.Context) {
    idStr := c.Param("id")
    id64, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        c.JSON(400, gin.H{"error": "invalid id"})
        return
    }

    var input models.EditShow
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": "invalid input"})
        return
    }

    updated, err := services.EditShow(uint(id64), &input)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"show": updated})
}

func DeleteShow(c *gin.Context){
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	er := services.DeleteShow(uint(id))
	if er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"show deleting faild",
			"err": er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"deleted successfull",
	})
}
