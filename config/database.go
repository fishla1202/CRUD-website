package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang_side_project_crud_website/models/collections"
	"golang_side_project_crud_website/models/posts"
	"golang_side_project_crud_website/models/users"
	"log"
	"os"
	"time"
)


func OpenDatabaseConnectionPool() *gorm.DB{

	dbConnectionKeyword := os.Getenv("DB_CONNECTION_KEYWORD")
	db, err := gorm.Open(
		"mysql", dbConnectionKeyword)
	if err != nil { log.Fatal("connection error:", err) }
	db.LogMode(true)
	db.DB()
	err = db.DB().Ping()
	if err != nil {log.Fatal("connection error:", err) }
	db.DB().SetMaxIdleConns(1000)
	db.DB().SetMaxOpenConns(2000)
	db.DB().SetConnMaxLifetime(time.Hour)
	return db
}

func MigrateAndPassDatabaseConnectionToModels(db *gorm.DB) {
	posts.DB = db
	posts.InitPostTable()
	users.DB = db
	users.InitUserTable()
	collections.DB = db
	collections.InitCollectionTable()
}
