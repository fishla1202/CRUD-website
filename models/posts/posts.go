package posts

import (
	"fmt"
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

// TODO: show the query res
func FindAllPosts() *gorm.DB{
	var posts []Post
	res := DB.Find(&posts)
	fmt.Println(DB)
	return res
}