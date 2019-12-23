package posts

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title string `gorm:"not null"`
	Content string `sql:"type:text;"gorm:"not null"`
	UserID  uint
}

var DB *gorm.DB

func InitPostTable() {
	DB.AutoMigrate(&Post{})
}

func InsertPost(postContent *Post) {
	DB.Create(postContent)
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

func FindByUserId(userID uint) *gorm.DB{
	var posts []Post
	res := DB.Find(&posts, Post{UserID: userID})
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