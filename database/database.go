package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"learnFiberAPI/models"
	"log"
	"os"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {

	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening connection to api.db", err.Error())
		os.Exit(2)
	}
	log.Println("Connected sucessfully to Database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	if err != nil {
		return
	}
	Database = DbInstance{Db: db}
}
