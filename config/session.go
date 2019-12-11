package config

import (
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"path"
	"time"
	//"firebase.google.com/go/auth"
	"golang.org/x/net/context"
	"os"
)

func SetLoginSession(w http.ResponseWriter, r *http.Request) {
	// Get the ID token sent by the client
	idToken, err := getIDTokenFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set session expiration to 5 days.
	expiresIn := time.Hour * 24 * 5

	// Create the session cookie. This will also verify the ID token in the process.
	// The session cookie will have the same claims as the ID token.
	// To only allow session cookie setting on recent sign-in, auth_time in ID token
	// can be checked to ensure user was recently signed in before creating a session cookie.
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

	//decoded, err := client.VerifyIDToken(r.Context(), idToken)
	//if err != nil {
	//	http.Error(w, "Invalid ID token", http.StatusUnauthorized)
	//	return
	//}

	// Return error if the sign-in is older than 5 minutes.
	//if time.Now().Unix()-decoded.Claims["auth_time"].(int64) > 5*60 {
	//	http.Error(w, "Recent sign-in required", http.StatusUnauthorized)
	//	return
	//}



	cookie, err := client.SessionCookie(r.Context(), idToken, expiresIn)
	if err != nil {
		http.Error(w, "Failed to create a session cookie", http.StatusInternalServerError)
		return
	}

	// Set cookie policy for session cookie.
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    cookie,
		MaxAge:   int(expiresIn.Seconds()),
		HttpOnly: true,
		Secure:   true,
	})
	w.Write([]byte(`{"status": "success"}`))
}

func getIDTokenFromBody(r *http.Request) (string, error){
	err := r.ParseForm()
	if err != nil {return "", err}
	return r.Form["idToken"][0], nil
}