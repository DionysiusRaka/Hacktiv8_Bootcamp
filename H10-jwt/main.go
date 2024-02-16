package main

import (
	"Hactiv8_Bootcamp/H10-jwt/database"
	"Hactiv8_Bootcamp/H10-jwt/routers"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8080")
}
