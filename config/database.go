package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var db *gorm.DB

func DatabaseInit() *gorm.DB{
	db, err := gorm.Open(
		"mysql", "root:root@tcp(127.0.0.1:3306)/golang?charset=utf8")
	if err != nil { log.Fatal("mysql connection error:", err)}

	db.DB()
	db.DB().SetMaxIdleConns(1000)
	db.DB().SetMaxOpenConns(2000)
	db.DB().SetConnMaxLifetime(time.Hour)
	return db
}
