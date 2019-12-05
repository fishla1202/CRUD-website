package posts

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Post struct {
	ID int64 `gorm:"PRIMARY_KEY"`
	Title string `gorm:"not null"`
	Content string `sql:"type:text;"gorm:"not null"`
	UserID  uint
	UpdatedAt *time.Time
	CreatedAt *time.Time
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
	var postsArray []Post
	res := DB.Find(&postsArray)
	//fmt.Println(res.Value)
	return res
}

func FindById(id string) *gorm.DB{
	var posts Post
	res := DB.Find(&posts, id)
	return res
}

func UpdateById(id string, title string, content string){
	var posts Post
	DB.Model(&posts).Where("id = ?", id).Update(
		map[string]interface{}{"title": title, "content": content})
}

func DeleteById(id string) {
	var posts Post
	DB.Where("id = ?", id).Delete(&posts)
}