package controllers

import (
	"net/http"

	"github.com/trevrosen/git-grabber/db"
)

// GitHubUserCreateUserType represents the payload coming off the wire
type GitHubUserCreateUserType struct {
	// Username is the GitHub username
	Username string
}

type GitHubUserMediaTyoe struct {
	// ID is the primary key in this database
	ID int
	// Username is the GitHub username
	Username string
}

// userCreateHandler handles POST /users/
func userCreateHandler(dbi db.DBInteractor) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create a user

	})
}

// userShowHandler handles GET /users/{username}
func userShowHandler(dbi db.DBInteractor) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// find/serialize a user
	})
}
