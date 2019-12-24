package models

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"os"
	"path"
	"time"
)

type User struct {
	ID uint `gorm:"PRIMARY_KEY"`
	Name string `gorm:"not null"`
	Email string `gorm:"not null;index:email_idx"`
	Uid string `gorm:"unique;not null;index:uid_idx"`
	Posts []Post
	UpdatedAt *time.Time
	CreatedAt *time.Time
}

var DB *gorm.DB

func InitUserTable() {
	DB.AutoMigrate(&User{})
}

func InsertUser(name string, email string, uid string) bool{
	user := User{Name: name, Email: email, Uid: uid}
	res := DB.Create(&user)
	if res.Error != nil {
		return false
	}
	return true
}

func CreateUser(userName string, userEmail string, userPwd string) (string, bool){
	opt := option.WithCredentialsFile(path.Join("secret", os.Getenv("FIREBASE_KEY_PATH")))
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		errMessage := fmt.Sprintf("error initializing app: %v\n", err)
		return errMessage, false
	}
	fmt.Println(app)

	client, err := app.Auth(ctx)
	if err != nil {
		errMessage := fmt.Sprintf("error creating user: %v\n", err)
		return errMessage, false
	}

	params := (&auth.UserToCreate{}).
		Email(userEmail).
		EmailVerified(false).
		Password(userPwd).
		DisplayName(userName).
		Disabled(false)

	u, err := client.CreateUser(ctx, params)
	if err != nil {
		errMessage := fmt.Sprintf("error creating user: %v\n", err)
		return errMessage, false
	}
	log.Printf("Successfully created user: %v\n", u)
	isCreated := InsertUser(userName, userEmail, u.UID)

	if !isCreated{
		return "Fail to create user for db", false
	}

	return u.UID, true
}

func FindUserByID(id string) (User, error){

	var user User
	return user, DB.Find(&user, id).Error
}


func FindUserIDByUID(uid string) (uint, error){
	var user User
	userId := DB.Select("id").Find(&user, User{Uid: uid})
	if userId.Error != nil {
		return user.ID, userId.Error
	}
	return user.ID, nil
}


