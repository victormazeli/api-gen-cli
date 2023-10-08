package handlers

import (
	"goApiStartetProject/internal/config"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Db  *gorm.DB
	Env *config.Env
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	// Handle user creation logic
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func (u *UserHandler) GetUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	// Fetch user from database by userID
	c.JSON(http.StatusOK, gin.H{"user_id": userID, "name": "John Doe"})
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	_, _ = strconv.Atoi(c.Param("id"))
	// Update user in the database by userID
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	_, _ = strconv.Atoi(c.Param("id"))
	// Delete user from the database by userID
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
