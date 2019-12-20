package config

import (
	"github.com/gorilla/sessions"
	"golang_side_project_crud_website/models/users"
	"net/http"
	"time"
)

var Store = sessions.NewCookieStore(make([]byte, 32))

func SetLoginSession(w http.ResponseWriter, r *http.Request) {
	// Get the ID token sent by the client
	defer r.Body.Close()
	userLoginInfo, err := getIDTokenFromBody(r)
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

	decoded, err := Client.VerifyIDToken(r.Context(), userLoginInfo["idToken"])
	if err != nil {
		http.Error(w, "Invalid ID token", http.StatusUnauthorized)
		return
	}

	//Return error if the sign-in is older than 5 minutes.
	if  float64(time.Now().Unix())-decoded.Claims["auth_time"].(float64) > 5*60 {
		http.Error(w, "Recent sign-in required", http.StatusUnauthorized)
		return
	}


	cookie, err := Client.SessionCookie(r.Context(), userLoginInfo["idToken"], expiresIn)
	if err != nil {
		http.Error(w, "Failed to create a session cookie", http.StatusInternalServerError)
		return
	}

	userId := users.FindUserByUID(userLoginInfo["uid"])
	session, _ := Store.Get(r, "user-info")
	// Set user session values.
	session.Values["uid"] = userLoginInfo["uid"]
	session.Values["sessionCookie"] = cookie
	session.Values["userId"] = userId
	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"status": "success"}`))
}

func getIDTokenFromBody(r *http.Request) (map[string]string, error){
	err := r.ParseForm()
	if err != nil {return nil, err}
	userLoginInfo := map[string]string {}
	userLoginInfo["idToken"] = r.Form["idToken"][0]
	userLoginInfo["uid"] = r.Form["uid"][0]
	return userLoginInfo, nil
}

func cleanSession(w http.ResponseWriter, r *http.Request){

	session, _ := Store.Get(r, "user-info")
	session.Values["sessionCookie"] = nil
	session.Values["uid"] = nil
	session.Values["userId"] = nil
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CheckSessionCookie(w http.ResponseWriter, r *http.Request) bool{
	defer r.Body.Close()

	// Get the ID token sent by the client
	session, _ := Store.Get(r, "user-info")

	if session.Values["sessionCookie"] == nil {
		cleanSession(w, r)
		return false
	}
	// Verify the session cookie. In this case an additional check is added to detect
	// if the user's Firebase session was revoked, user deleted/disabled, etc.
	_, err := Client.VerifySessionCookieAndCheckRevoked(r.Context(), session.Values["sessionCookie"].(string))
	if err != nil {
		// Session cookie is invalid. Force user to login.
		cleanSession(w, r)
		return false
	}
	return true
}

func SessionSignOut(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	session, err := Store.Get(r, "user-info")
	if err != nil {
		// Session cookie is unavailable. Force user to login.
		http.Redirect(w, r, "/user/login/", http.StatusFound)
		return
	}

	decoded, err := Client.VerifySessionCookie(r.Context(), session.Values["sessionCookie"].(string))
	if err != nil {
		// Session cookie is invalid. Force user to login.
		http.Redirect(w, r, "/user/login/", http.StatusFound)
		return
	}
	if err := Client.RevokeRefreshTokens(r.Context(), decoded.UID); err != nil {
		http.Error(w, "Failed to revoke refresh token", http.StatusInternalServerError)
		return
	}

	cleanSession(w, r)

	http.Redirect(w, r, "/", http.StatusFound)
}