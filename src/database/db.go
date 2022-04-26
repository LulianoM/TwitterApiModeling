package database

import (
	"apiposterr/src/structs"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	stringDeConexao := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(structs.Post{}, structs.User{})
}
