package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title      string `gorm:"not null"`
	Content    string `sql:"type:text;"gorm:"not null"`
	UserID     uint	`gorm:"not null"`
	CollectionID     uint `gorm:"not null"`
	User User
	Collection Collection
	Comments []Comments
}


func InitPostTable() {
	db.AutoMigrate(&Post{})
}

func InsertPost(postContent *Post) error{
	return db.Create(postContent).Error
}

func FindAllPosts() ([]Post, error){
	var postsArray []Post
	return postsArray, db.Preload("Collection").Preload("User").Find(&postsArray).Error
}

func FindPostById(id string) (Post, error){
	var post Post
	return post, db.Preload("Collection").Preload("Comments").Preload("Comments.User").Find(&post, id).Error
}

func FindPostByCollectionTitle(collectionTitle string) ([]Post, error) {
	var posts []Post
	collection, err := FindCollectionByTitle(collectionTitle)
	if err != nil {
		return posts, err
	}
	return posts, db.Preload("User").Find(&posts, Post{CollectionID: collection.ID}).Error
}

func FindPostByUserId(userID uint) ([]Post, error){
	var post []Post
	return post, db.Find(&post, Post{UserID: userID}).Error
}

func UpdatePostById(id string, updatePostMap map[string]interface{}) error{
	var post Post
	return db.Model(&post).Where("id = ?", id).Update(updatePostMap).Error
}

func DeletePostById(id string) error{
	var post Post
	return db.Where("id = ?", id).Delete(&post).Error
}