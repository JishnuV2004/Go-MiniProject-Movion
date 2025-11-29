package controllers

import (
	services "movion/Services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBooking(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
		return
	}
	userID := userIDVal.(uint)

	var input struct {
		ShowID uint     `json:"show_id"`
		Seats  []string `json:"seats"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	booking, err := services.BookSeats(userID, input.ShowID, input.Seats)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"booking": booking})
}
func CancelBooking(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
		return
	}
	userID := userIDVal.(uint)

	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := services.CancelBooking(uint(id64), userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "booking cancelled"})
}
