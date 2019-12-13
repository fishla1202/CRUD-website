package config

import (
	"firebase.google.com/go/auth"
	"net/http"
	"time"
)

func SetLoginSession(w http.ResponseWriter, r *http.Request) {
	// Get the ID token sent by the client
	defer r.Body.Close()
	idToken, err := getIDTokenFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set session expiration to 1 days.
	expiresIn := time.Hour * 24

	// Create the session cookie. This will also verify the ID token in the process.
	// The session cookie will have the same claims as the ID token.
	// To only allow session cookie setting on recent sign-in, auth_time in ID token
	// can be checked to ensure user was recently signed in before creating a session cookie.
	client := GetFireBaseClient()

	decoded, err := client.VerifyIDToken(r.Context(), idToken)
	if err != nil {
		http.Error(w, "Invalid ID token", http.StatusUnauthorized)
		return
	}

	//Return error if the sign-in is older than 5 minutes.
	if  float64(time.Now().Unix())-decoded.Claims["auth_time"].(float64) > 5*60 {
		http.Error(w, "Recent sign-in required", http.StatusUnauthorized)
		return
	}


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

func CheackSessionCookie(w http.ResponseWriter, r *http.Request) *auth.Token{
	// Get the ID token sent by the client
	cookie, err := r.Cookie("session")
	//fmt.Print(cookie)
	if err != nil {
		// Session cookie is unavailable. Force user to login.
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	client := GetFireBaseClient()
	// Verify the session cookie. In this case an additional check is added to detect
	// if the user's Firebase session was revoked, user deleted/disabled, etc.
	decoded, err := client.VerifySessionCookieAndCheckRevoked(r.Context(), cookie.Value)
	if err != nil {
		// Session cookie is invalid. Force user to login.
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	return decoded
}