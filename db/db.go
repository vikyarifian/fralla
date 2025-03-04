package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var PgSql *gorm.DB

func ConnectDB() {

	var err error
	dsn := os.Getenv("POSTGRES_URL")
	PgSql, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

}

func ConnectDBVercel() {

	var err error
	// dsn := "postgres://postgres.akqzqmwrlvegatfgbweu:QPIjWeXg6GT2mHGK@aws-0-ap-southeast-1.pooler.supabase.com:6543/db_fralla?prepareThreshold=0&sslmode=require&supa=base-pooler.x"
	dsn := "postgres://neondb_owner:npg_KRwc1WXE5iJY@ep-super-heart-a14v5sbg-pooler.ap-southeast-1.aws.neon.tech/db_fralla"
	PgSql, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

}
