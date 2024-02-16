package main

import (
	"Hactiv8_Bootcamp/H8-gin/database"
	"Hactiv8_Bootcamp/H8-gin/routers"
)

func main() {
	var PORT = ":8080"
	database.StartDB()
	routers.StartServer().Run(PORT)
}
