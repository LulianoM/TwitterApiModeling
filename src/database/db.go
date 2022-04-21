package database

import (
	"apiposterr/src/structs"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnection() {

	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Could not connect with the database!")
	}
}

func DatabaseMigration() {
	DB.AutoMigrate(structs.User{})
}
