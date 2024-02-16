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

func CreateSocMed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocMed := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&SocMed)
	} else {
		c.ShouldBind(&SocMed)
	}

	SocMed.UserId = userID

	err := db.Debug().Create(&SocMed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocMed)
}

func GetAllSocMed(c *gin.Context) {
	db := database.GetDB()
	SocMed := []models.SocialMedia{}

	err := db.Find(&SocMed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocMed)
}

func GetSocMedByID(c *gin.Context) {
	db := database.GetDB()

	SocMedID, err := strconv.Atoi(c.Param("socMedId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Bad Request",
			"error_message": "Social Media not found",
		})
		return
	}

	socMed := models.SocialMedia{}
	err = db.First(&socMed, "id=?", SocMedID).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Social Media not found",
			"error_message": "Social Media not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comment": socMed,
	})
}

func UpdateSocMed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocMed := models.SocialMedia{}

	SocMedId, _ := strconv.Atoi(c.Param("socMedId"))
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&SocMed)
	} else {
		c.ShouldBind(&SocMed)
	}

	SocMed.UserId = userID
	SocMed.ID = uint(SocMedId)

	err := db.Model(&SocMed).Where("id = ?", SocMedId).Updates(models.SocialMedia{Name: SocMed.Name, SocialMediaURL: SocMed.SocialMediaURL}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocMed)
}

func DeleteSocMed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocMed := models.SocialMedia{}

	SocMedID, _ := strconv.Atoi(c.Param("socMedId"))
	userID := uint(userData["id"].(float64))

	SocMed.UserId = userID
	SocMed.ID = uint(SocMedID)

	err := db.First(&SocMed, "id = ?", SocMedID).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	db.Delete(&SocMed)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Social media deleted successfully",
	})
}
