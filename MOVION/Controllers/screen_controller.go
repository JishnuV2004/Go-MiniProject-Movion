package controllers

import (
	models "movion/Models"
	services "movion/Services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
// POST /admin/screens?rows=10&cols=10
func CreateScreen(c *gin.Context){
	var screenInput models.Screen
	if err := c.ShouldBindJSON(&screenInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid request",
		})
		return
	}
	rowsStr := c.DefaultQuery("row", "10")
	colsStr := c.DefaultQuery("col", "10")
	row,_ := strconv.Atoi(rowsStr)
	col,_ := strconv.Atoi(colsStr)

	screen, err := services.CreateScreen(&screenInput, row, col)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"screen creation failed",
			"err":err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"screen":screen,
	})
}

func GetAllScreens(c *gin.Context){
	screens, err := services.GetAllScreens()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"screens getting failed",
			"err":err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"screens":screens,
	})
}
func GetScreen(c *gin.Context){
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid request",
		})
		return
	}
	screen, err := services.GetScreen(uint(id64))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"screen not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"screen":screen,
	})
}
func DeleteScreen(c *gin.Context){
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid request",
		})
		return
	}
	 er := services.DeleteScreen(uint(id64)); 
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"deleting failed",
			"err":er.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"deleting successful",
	})
}
func EditScreen(c *gin.Context){
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var input models.Screen
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid input",
		})
		return
	}
	editedscreen, err := services.EditScreen(uint(id64), &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"editng faild",
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"edit successful",
		"screen": editedscreen,
	})
}