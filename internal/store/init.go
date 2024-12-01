package store

import (
	"authService/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"

	"log"
)

var Db *gorm.DB

func DataBaseInit() {
	host := os.Getenv("DB_HOST")
	port := 5432
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s ",
		host, port, user, password, dbname)
	log.Println(dsn)
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	err = Db.AutoMigrate(&models.User{})
	if err != nil {
		log.Println(err)
	}
}
