package posts

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title string `gorm:"not null"`
	Content string `sql:"type:text;"gorm:"not null"`
}

var DB *gorm.DB

func InitPostTable() {
	DB.AutoMigrate(&Post{})
}

func InsertPost(title string, content string) {
	post := Post{Title: title, Content: content}
	DB.Create(&post)
}

func FindAllPosts() *gorm.DB{
	res := DB.Find(&[]Post{})
	return res
}