package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang_side_project_crud_website/models/posts"
	"log"
	"time"
)


func OpenDatabaseConnectionPool() *gorm.DB{
	// TODO: use os.env
	db, err := gorm.Open(
		"mysql", "root:root@tcp(127.0.0.1:3306)/golang?charset=utf8")
	if err != nil { log.Fatal("connection error:", err) }

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
}
