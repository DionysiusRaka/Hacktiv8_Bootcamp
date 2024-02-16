package database

import (
	"Hactiv8_Bootcamp/H10-jwt/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Host   = "localhost"
	User   = "postgres"
	Pass   = "password"
	Port   = "5432"
	Dbname = "simple-api"
	Db     *gorm.DB
	Err    error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", Host, User, Pass, Dbname, Port)
	dsn := config
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database : ", err)
	}

	fmt.Println("Database connected")
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=simple-api port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database : ", err)
	}

	fmt.Println("Database connected")
	db.Debug().AutoMigrate(models.User{}, models.Product{})
	return db
}
