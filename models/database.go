package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func init() {

	dbConnectionKeyword := os.Getenv("DB_CONNECTION_KEYWORD")
	DB, err := gorm.Open(
		"mysql", dbConnectionKeyword)
	if err != nil { log.Fatal("connection error:", err) }
	DB.LogMode(true)
	DB.DB()
	err = DB.DB().Ping()
	if err != nil {log.Fatal("connection error:", err) }
	DB.DB().SetMaxIdleConns(1000)
	DB.DB().SetMaxOpenConns(2000)
	DB.DB().SetConnMaxLifetime(time.Hour)
	db = DB
	migrateAndPassDatabaseConnectionToModels()
}

func migrateAndPassDatabaseConnectionToModels() {
	InitPostTable()
	InitUserTable()
	InitCollectionTable()
	InitCommentTable()
}