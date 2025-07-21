package handler

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"vue-gin-backend/internal/model"
	"vue-gin-backend/pkg/database"
)

func GetUsers(c *gin.Context) {
	var users []model.User
	
	result := database.DB.Preload("Microposts").Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
		})
		return
	}
	
	c.JSON(http.StatusOK, users)
}
