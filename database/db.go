package database

import (
	"DTSGolang/FinalProject/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "1234"
	dbName   = "dtsgo-final-project"
	dbPort   = 5432
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
	}

	fmt.Println("Successfully connected to the database")
	db.Debug().AutoMigrate(models.User{}, &models.SocialMedia{}, &models.Photo{}, &models.Comment{})
}

func CloseDB() {
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	if sqlDb != nil {
		err := sqlDb.Close()
		if err != nil {
			log.Fatal("error closing to database: ", err.Error())
		}
	}
}

func GetDB() *gorm.DB {
	return db
}
