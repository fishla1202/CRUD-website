package models

import "time"

type Comments struct {
	ID 		uint `gorm:"PRIMARY_KEY"`
	Content	string `sql:"type:text;"gorm:"not null"`
	UserID	uint
	User	User
	PostID 	uint
	Posts 	Post
	UpdatedAt *time.Time
	CreatedAt *time.Time
}


func InitCommentTable() {
	DB.AutoMigrate(&Post{})
}