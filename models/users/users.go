package users

// TODO: add firebase sign up with email password

// https://firebase.google.com/docs/admin/setup?authuser=1
// https://firebase.google.com/docs/auth/admin/manage-users?authuser=1

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"path"

	//"firebase.google.com/go/auth"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"os"
)

func CreateUser(userName string, userEmail string, userPwd string) string{
	opt := option.WithCredentialsFile(path.Join(os.Getenv("FIREBASE_KEY_PATH")))
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
	return u.UID
}


