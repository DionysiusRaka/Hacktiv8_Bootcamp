package controllers

import (
	"Hactiv8_Bootcamp/final-project-mygram/database"
	"Hactiv8_Bootcamp/final-project-mygram/helpers"
	"Hactiv8_Bootcamp/final-project-mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	newComment := models.Comment{}
	userId := uint(userData["id"].(float64))
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	if contentType == "application/json" {
		c.ShouldBindJSON(&newComment)
	} else {
		c.ShouldBind(&newComment)
	}

	newComment.UserId = userId
	newComment.PhotoId = uint(photoId)
	err := db.Debug().Create(&newComment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_status":  "Bad Request",
			"error_message": "Bad Request",
		})
		return
	}

	c.JSON(http.StatusCreated, newComment)
}

func GetAllComments(c *gin.Context) {
	db := database.GetDB()
	Comments := []models.Comment{}

	err := db.Find(&Comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comments)
}

func GetCommentByID(c *gin.Context) {
	db := database.GetDB()

	commentID, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Bad Request",
			"error_message": "Comment not found",
		})
		return
	}

	comment := models.Comment{}
	err = db.Preload("User").Preload("Photo").First(&comment, commentID).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Comment not found",
			"error_message": "Comment not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comment": comment,
	})
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserId = userID
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comment := models.Comment{}

	commentID, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	Comment.UserId = userID
	Comment.ID = uint(commentID)

	err := db.First(&Comment, "id = ?", commentID).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	db.Delete(&Comment)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Comment deleted successfully",
	})
}
