package api

import (
	"Teamwork-Golang/creating"
	"Teamwork-Golang/deleting"
	"Teamwork-Golang/updating"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateUser(service creating.CreatingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := creating.Users{}
		// read the request body into byte slice
		requestBody, _ := ioutil.ReadAll(r.Body)
		// parse the body into the user slice, handle any errors
		if err := json.Unmarshal(requestBody, &user); err != nil {
			writeError(w, err, string(requestBody))
			return
		}
		// send user slice to the creating service, handle any errors
		UserID, err := service.CreateUser(user)
		if err != nil {
			writeError(w, err, user)
			return
		}

		// write a success response back to the originating service
		var response = operationResponse{"SUCCESS", "user account successfully", UserID, make([]string, 0)}
		writeJSON(w, response)
	}
}

func UserSignIn(service creating.CreatingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userinfo := creating.UserSignInfo{}
		err := json.NewDecoder(r.Body).Decode(&userinfo)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		erro := service.UserSignIn(userinfo)
		if erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

func CreateArticle(service creating.CreatingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		article := creating.Article{}
		err := json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		articleID, timeCreated, articleTitle, erro := service.CreateArticle(article)
		if erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var response = articleResponse{"Success", "Article successfully posted", articleID, timeCreated, articleTitle, make([]string, 0)}
		writeJSON(w, response)
	}
}

func UpdateArticle(service updating.UpdateService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		update := updating.UpdateAtricle{}
		err := json.NewDecoder(r.Body).Decode(&update)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		articleTitle, message, err := service.UpdateArticle(update)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var resp = articleUpdateResponse{"success", message, articleTitle, make([]string, 0)}
		writeJSON(w, resp)

	}
}

func DeleteArticle(service deleting.DeleteService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		delete := deleting.DeleteArt{}
		err := json.NewDecoder(r.Body).Decode(&delete)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		erro := service.DeleteArticle(delete)
		if erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var resp = deletingUpdateResponse{"success", "article successfully deleted"}
		writeJSON(w, resp)

	}
}
