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
	db.AutoMigrate(&Collection{})
}

func CreateCollection(collection *Collection) error{
	return db.Create(&collection).Error
}

func FindAllCollections() ([]Collection, error){
	var collections []Collection
	return collections, db.Find(&collections).Error
}

func FindCollectionByID(id string) (Collection, error){
	var collection Collection
	return collection, db.Find(&collection, id).Error
}

func FindCollectionByTitle(title string) (Collection, error) {
	var collection Collection
	return collection, db.Find(&collection, Collection{Title: title}).Error
}
