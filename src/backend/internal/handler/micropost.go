package handler

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"vue-gin-backend/internal/model"
	"vue-gin-backend/pkg/database"
)

func GetMicroposts(c *gin.Context) {
	var microposts []model.Micropost
	
	result := database.DB.Preload("User").Find(&microposts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch microposts",
		})
		return
	}
	
	c.JSON(http.StatusOK, microposts)
}

func GetMicropostByID(c *gin.Context) {
	id := c.Param("id")
	
	var micropost model.Micropost
	result := database.DB.Preload("User").First(&micropost, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Micropost not found",
		})
		return
	}
	
	c.JSON(http.StatusOK, micropost)
}

type CreateMicropostRequest struct {
	Title  string `json:"title" binding:"required,max=200"`
	UserID uint   `json:"userId" binding:"required"`
}

func CreateMicropost(c *gin.Context) {
	var req CreateMicropostRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}
	
	var user model.User
	if err := database.DB.First(&user, req.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}
	
	micropost := model.Micropost{
		Title:  req.Title,
		UserID: req.UserID,
	}
	
	result := database.DB.Create(&micropost)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create micropost",
		})
		return
	}
	
	database.DB.Preload("User").First(&micropost, micropost.ID)
	
	c.JSON(http.StatusCreated, micropost)
}
