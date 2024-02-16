package main

import (
	"Hactiv8_Bootcamp/final-project-mygram/database"
	"Hactiv8_Bootcamp/final-project-mygram/routers"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8080")
}
