package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title      string `gorm:"not null"`
	Content    string `sql:"type:text;"gorm:"not null"`
	UserID     uint
	CollectionID     uint
	Collection Collection
	Comments []Comments
}


func InitPostTable() {
	DB.AutoMigrate(&Post{})
}

func InsertPost(postContent *Post) error{
	return DB.Create(postContent).Error
}

func FindAllPosts() ([]Post, error){
	var postsArray []Post
	return postsArray, DB.Preload("Collection").Find(&postsArray).Error
}

func FindPostById(id string) (Post, error){
	var post Post
	return post, DB.Preload("Collection").Find(&post, id).Error
}

func FindPostByUserId(userID uint) ([]Post, error){
	var post []Post
	return post, DB.Find(&post, Post{UserID: userID}).Error
}

func UpdatePostById(id string, updatePostMap map[string]interface{}) error{
	var post Post
	return DB.Model(&post).Where("id = ?", id).Update(updatePostMap).Error
}

func DeletePostById(id string) error{
	var post Post
	return DB.Where("id = ?", id).Delete(&post).Error
}