package controllers

import (
	models "movion/Models"
	services "movion/Services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Signup func
func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	err := services.Signup(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "signup successful",
	})
}

// Login func
func Login(c *gin.Context) {
	var loginUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	accesstoken, refreshtoken, err := services.Login(loginUser.Email, loginUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.SetCookie("refresh_token", refreshtoken, int(24*time.Hour.Seconds()), "/", "localhost", true, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		// "user":user.Username,
		"accessToken": accesstoken,
		"refresh":     refreshtoken,
	})
}

// Logout func
func Logout(c *gin.Context) {
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "logout successful",
	})
}
