package controllers

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/trevrosen/git-grabber/db"
	"github.com/trevrosen/git-grabber/middleware"
	"github.com/urfave/negroni"
)

func App(dbi db.DBInteractor) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/status", statusHandler()).Methods("GET")
	router.HandleFunc("/users", userCreateHandler(dbi)).Methods("POST")
	router.HandleFunc("/users/{username}", userShowHandler(dbi)).Methods("GET")

	n := negroni.New(negroni.NewRecovery())
	n.Use(middleware.NewGGLogger())    // Log with Logrus
	n.Use(middleware.NewContentType()) // Ensure response Content-Type header is always "application/json"
	n.UseHandler(router)
	return n
}

// statusHandler handles GET /status
// provides system status information
func statusHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Up and running"))
		logMsg(r, "All systems nominal, Captain")
	})
}

// logMsg sets a logging message on the middleware.HttpInformer that is set up for each request
// by the logging middleware
func logMsg(r *http.Request, msg string) {
	httpInformant := context.Get(r, middleware.HttpInformantKey)
	hi := httpInformant.(*middleware.HttpInformant)
	hi.SetMessage(msg)
}
