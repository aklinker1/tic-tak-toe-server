package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(entities ...interface{}) *gorm.DB {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		panic("DATABASE_URL is required in the env")
	}

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(entities...)
	if err != nil {
		panic(err)
	}

	return db
}
