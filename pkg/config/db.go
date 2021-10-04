package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func Connect() {
	godotenv.Load()
	d, err := gorm.Open("postgres", os.Getenv("DB_URI"))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connected to database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
