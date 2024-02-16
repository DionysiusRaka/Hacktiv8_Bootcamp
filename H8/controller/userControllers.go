package controller

import (
	"Hactiv8_Bootcamp/H8/database"
	"Hactiv8_Bootcamp/H8/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CreateUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data : ", err)
		return
	}

	fmt.Println("New user data : ", User)
}

func GetUserbyId(id uint) {
	db := database.GetDB()

	User := models.User{}

	err := db.First(&User, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user : ", err)
	}
	fmt.Printf("User Data : %v \n ", User)
}

func GetAllUsers() {
	db := database.GetDB()

	Users := []models.User{}

	err := db.Find(&Users).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user : ", err)
	}
	for i, Users := range Users {
		fmt.Printf("User Data ke %v : %v \n ", i, Users)
	}
}

func UpdateUserById(id uint, email string) {
	db := database.GetDB()
	Users := models.User{}

	err := db.Model(&Users).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updating user data : ", err)
		return
	}
	fmt.Printf("Update user's email : %v \n ", Users.Email)
}
