package controllers

import (
	models "movion/Models"
	services "movion/Services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context){
	var adminlogin models.Admin
	if err := c.ShouldBindJSON(&adminlogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid request",
			"err":err.Error(),
		})
		return
	}
	admin, err := services.AdminLogin(adminlogin.Email, adminlogin.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":"user not found",
			"err": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"admin": admin.Email,
		"message":"login successful",
	})
}
// Get all users
func GetAllUsers(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	users, err := services.GetAllUsers(page, limit)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "users not found",
		})
		return
	}

	c.HTML(http.StatusOK, "layout.html", gin.H{
    "PageTitle": "Users",
    "Users": users,
})

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// Get user
func GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	user, err := services.GetUser(id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"user":  user.Username,
		"email": user.Email,
	})
}

// update user
func EditUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invaied id",
		})
		return
	}

	var input models.EditUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	updatedUser, err := services.EditUser(id, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// avoid returning password field in response
	updatedUser.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"message": "edit successful",
		"user":    updatedUser,
	})
}

// Create users
func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}
	user, err := services.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user creating successful",
		"user":    user,
	})
}

// Delete users
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	if err := services.DeleteUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "delete successful",
	})
}

// Search users
func SearchUser(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name is requerid",
		})
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	users, err := services.SearchUser(name, page, limit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": " successful",
		"users":   users,
	})
}

// PUT /admin/users/:id/block?block=true
func BlockUser(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	blockStr := c.DefaultQuery("block", "true")
	block := true
	if blockStr == "false" {
		block = false
	}
	user, err := services.BlockUser(id, block)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
