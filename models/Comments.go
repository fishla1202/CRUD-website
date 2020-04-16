package models

import "time"

type Comments struct {
	ID 		uint `gorm:"PRIMARY_KEY"`
	Content	string `sql:"type:text;"gorm:"not null"`
	LikeCount	int64
	UserID	uint `gorm:"not null"`
	User	User
	PostID 	uint `gorm:"not null"`
	Post 	Post
	UpdatedAt *time.Time
	CreatedAt *time.Time
}


func InitCommentTable() {
	db.AutoMigrate(&Comments{})
}

func CreateComment(comment *Comments) error{
	return db.Create(comment).Error
}