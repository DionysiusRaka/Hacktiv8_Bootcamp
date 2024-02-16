package main

import (
	"Hactiv8_Bootcamp/H8/controller"
	"Hactiv8_Bootcamp/H8/database"
)

func main() {
	database.StartDB()
	// controller.CreateUser("the.galonisky@gmail.com")
	controller.GetUserbyId(1)
	// controller.GetAllUsers()
	// controller.UpdateUserById(1, "microsoftex@email.com")
}
