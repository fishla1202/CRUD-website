package migrations

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Title string
	Content string
}

func InitPostTable(db *gorm.DB) {
	db.AutoMigrate(&Post{})
}