package config

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"os"
	"path"
)

var Client *auth.Client

func InitFirebaseClient() {
	opt := option.WithCredentialsFile(path.Join("secret", os.Getenv("FIREBASE_KEY_PATH")))
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatal("error initializing app:", err)
	}

	Client, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
}
