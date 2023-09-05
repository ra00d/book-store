package database

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	godotenv.Load()
	db, err := gorm.Open(sqlite.Open(os.Getenv("GOOSE_DBSTRING")), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
