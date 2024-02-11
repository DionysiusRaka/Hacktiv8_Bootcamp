package main

import (
	"Hactiv8_Bootcamp/H8-Assign/db"
	"Hactiv8_Bootcamp/H8-Assign/routers"
)

func main() {
	var PORT = ":8080"
	db.StartDB()
	routers.StartServer().Run(PORT)

}
