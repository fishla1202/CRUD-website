package users

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/jinzhu/gorm"
	"golang_side_project_crud_website/models/posts"
	"path"
	"time"

	//"firebase.google.com/go/auth"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"os"
)

type User struct {
	ID int64 `gorm:"PRIMARY_KEY"`
	Name string `gorm:"not null"`
	Email string `gorm:"not null;index:email_idx"`
	Uid string `gorm:"unique;not null;index:uid_idx"`
	Posts []posts.Post
	UpdatedAt *time.Time
	CreatedAt *time.Time
}

var DB *gorm.DB

func InitUserTable() {
	DB.AutoMigrate(&User{})
}

func InsertUser(name string, email string, uid string) {
	user := User{Name: name, Email: email, Uid: uid}
	DB.Create(&user)
}

func CreateUser(userName string, userEmail string, userPwd string) string{
	opt := option.WithCredentialsFile(path.Join("secret", os.Getenv("FIREBASE_KEY_PATH")))
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal("error initializing app:", err)}
	fmt.Println(app)

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}

	params := (&auth.UserToCreate{}).
		Email(userEmail).
		EmailVerified(false).
		Password(userPwd).
		DisplayName(userName).
		Disabled(false)
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %v\n", u)
	InsertUser(userName, userEmail, u.UID)
	return u.UID
}

func FindUserByUID(uid string) {
	//TODO: 查詢使用用者id 回傳
	//var user User
	//user := DB.Where("uid = ?", uid).Select("id").First(&user)
}


