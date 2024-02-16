package controllers

import (
	"Hactiv8_Bootcamp/H10-jwt/database"
	"Hactiv8_Bootcamp/H10-jwt/helpers"
	"Hactiv8_Bootcamp/H10-jwt/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	log.Println(&User)

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"full_name": User.FullName,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	fmt.Println(User.Email, " : email Raw")
	fmt.Println(&User.Email, " : email User")
	password = User.Password
	err := db.Debug().Where("email =?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password email 1",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	fmt.Println(password, " : Password Raw")
	fmt.Println(User.Password, " : Password User")

	fmt.Println(comparePass, " : Password Compare")
	fmt.Println(comparePass)
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password password 1",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
