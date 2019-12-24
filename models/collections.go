package models

import (
	"github.com/jinzhu/gorm"
)

type Collection struct {
	gorm.Model
	Title string `gorm:"not null"`
	Description string `sql:"type:text;"gorm:"not null"`
	Posts []Post
}


func InitCollectionTable() {
	DB.AutoMigrate(&Collection{})
}

func CreateCollection(collection *Collection) error{
	return DB.Create(&collection).Error
}

func FindAllCollections() ([]Collection, error){
	var collections []Collection
	return collections, DB.Find(&collections).Error
}

func FindCollectionByID(id string) (Collection, error){
	var collection Collection
	return collection, DB.Find(&collection, id).Error
}
