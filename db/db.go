package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PgSql *gorm.DB

func ConnectDB() {

	var err error
	dsn := os.Getenv("POSTGRES_URL")
	PgSql, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

}
