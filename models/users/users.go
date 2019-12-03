package users

// TODO: add firebase authentication
// https://firebase.google.com/docs/admin/setup?authuser=1
// https://firebase.google.com/docs/auth/admin/manage-users?authuser=1


import (
	firebase "firebase.google.com/go"
	//"firebase.google.com/go/auth"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"os"
)

func mian() {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_KEY_PATH"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("error initializing app:", err)}
	fmt.Println(app)
}


