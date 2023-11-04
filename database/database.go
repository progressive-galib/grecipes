package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"))
	if err != nil {
		log.Fatal("Failed to connect to DB")
		os.Exit(2)
	}
	log.Println("connected to db succesfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	//todo add migrations
	Database = DbInstance{Db: db}
}
