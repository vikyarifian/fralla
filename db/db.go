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

func ConnectDBVercel() {

	var err error
	dsn := "postgres://postgres.akqzqmwrlvegatfgbweu:QPIjWeXg6GT2mHGK@aws-0-ap-southeast-1.pooler.supabase.com:6543/db_fralla?sslmode=require&supa=base-pooler.x"
	PgSql, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

}
