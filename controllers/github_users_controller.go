package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/trevrosen/git-grabber/db"
)

// GitHubUserCreateUserType represents the payload coming off the wire
type GitHubUserCreateUserType struct {
	// Username is the GitHub username
	Username string
}

type GitHubUserMediaType struct {
	// ID is the primary key in this database
	ID uint
	// Username is the GitHub username
	Username string
}

// userCreateHandler handles POST /users/
func userCreateHandler(dbi db.DBInteractor) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		input := &GitHubUserCreateUserType{}
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(input)

		if err != nil {
			logMsg(r, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ghUser := &db.GitHubUser{
			Username: input.Username,
		}

		// TODO: better status code?
		if err = dbi.SaveRecord(ghUser); err != nil {
			logMsg(r, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// TODO: better status code?
		jsonBytes, err := serializedGitHubUser(ghUser)
		if err != nil {
			logMsg(r, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(jsonBytes)
		w.WriteHeader(http.StatusCreated)
	})
}

// userShowHandler handles GET /users/{username}
func userShowHandler(dbi db.DBInteractor) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		var ghUser *db.GitHubUser
		username := mux.Vars(r)["username"]

		if ghUser, err = dbi.FindGitHubUserByName(username); err != nil {
			if err == gorm.ErrRecordNotFound {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// TODO: better status code?
		jsonBytes, err := serializedGitHubUser(ghUser)
		if err != nil {
			logMsg(r, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(jsonBytes)
	})
}

// serializedGitHubUser serializes a db.GitHubUser for the wire
func serializedGitHubUser(ghUser *db.GitHubUser) ([]byte, error) {
	jsonSource := &GitHubUserMediaType{}

	jsonSource.ID = ghUser.ID
	jsonSource.Username = ghUser.Username

	return json.Marshal(jsonSource)
}
