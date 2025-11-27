package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBooking(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
		return
	}
	userID := userIDVal.(uint)

	var req struct {
		ShowID uint     `json:"show_id"`
		Seats  []string `json:"seats"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
}
