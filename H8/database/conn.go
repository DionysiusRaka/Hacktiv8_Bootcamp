package database

import (
	"Hactiv8_Bootcamp/H8/env"
	"Hactiv8_Bootcamp/H8/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", env.Host, env.User, env.Pass, env.Dbname, env.Port)

	env.Db, env.Err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if env.Err != nil {
		log.Fatal("error connecting to database : ", env.Err)
	}

	env.Db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return env.Db
}
