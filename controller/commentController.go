package controller

import (
	"encoding/json"
	"fmt"
	"golang_side_project_crud_website/config"
	"golang_side_project_crud_website/models"
	"net/http"
)


func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}


func CreateComment(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	isUser := config.CheckSessionCookie(w, r)

	if !isUser {
		responseWithJson(w, http.StatusBadRequest, "User not login")
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			responseWithJson(w, http.StatusInternalServerError, err.Error())
			return
		}

		if r.Form["content"] == nil || r.Form["postID"] == nil{
			responseWithJson(w, http.StatusBadRequest, "Missing required args")
			return
		}else {
			session, _ := config.Store.Get(r, "user-info")
			userId := session.Values["userId"]


			post, err := models.FindPostById(r.Form["postID"][0])

			if err != nil {
				responseWithJson(w, http.StatusInternalServerError, err.Error())
				return
			}
			fmt.Print("hi")
			user, err := models.FindUserByID(userId.(uint))
			if err != nil {
				responseWithJson(w, http.StatusInternalServerError, err.Error())
				return
			}

			comment := models.Comments{
				Content: r.Form["content"][0],
				UserID: userId.(uint),
				User: user,
				PostID: post.ID,
				Post: post,
			}

			err = models.CreateComment(&comment)
			if err != nil {
				responseWithJson(w, http.StatusInternalServerError, err.Error())
				return
			}

			responseWithJson(w, http.StatusCreated, "Comment is created")
			return
		}
	}else{
		responseWithJson(w, http.StatusMethodNotAllowed, "123")
	}
}