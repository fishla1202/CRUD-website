package collections

import (
	"github.com/jinzhu/gorm"
	"golang_side_project_crud_website/models/posts"
)

type Collection struct {
	gorm.Model
	Title string `gorm:"not null"`
	Description string `sql:"type:text;"gorm:"not null"`
	Posts []posts.Post
}

var DB *gorm.DB

func InitCollectionTable() {
	DB.AutoMigrate(&Collection{})
}

func CreateCollection(collection *Collection){
	DB.Create(&collection)
}