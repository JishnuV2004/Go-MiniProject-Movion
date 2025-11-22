package controllers

import (
	models "movion/Models"
	services "movion/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// user profile
func Profile(c *gin.Context){
	userID,_ := c.Get("userID")

	user, err := services.Profile(userID.(uint))

	if err  != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":"welcome to profile",
		"id":user.ID,
		"username":user.Username,
		"email":user.Email,
		"role":user.Role,
	})
}

// user update
func UpdateUser(c *gin.Context){
	userID,_:= c.Get("userID")

	if userID == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid request",
		})
		return
	}
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid input",
		})
		return
	}

	updatedUser, err := services.UpdateUser(userID.(uint), &input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message":"update success",
		"username": updatedUser.Username,
		"password":input.Password,
	})

}