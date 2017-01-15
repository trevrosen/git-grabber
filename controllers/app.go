package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func App(dbi db.DBInteractor) http.Handler {
	router := mux.NewRouter().StripSlash(true)

	router.HandleFunc("/username/{username}", usernameShowHandler()).Method("GET")

	// TODO: add logging middleware
}

func usernameShowHandler(dbi db.DBInteractor) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
