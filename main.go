package main

import (
	"fmt"
	"net/http"

	"github.com/bypasslane/alicia/middleware"
	"github.com/gorilla/context"
	"github.com/trevrosen/git-grabber/controllers"
)

func main() {
	// establish environment

	// allow override of environment via flag

	//

	bindAndRun()
}

// bindAndRun starts the server
func bindAndRun() {
	//portString := fmt.Sprintf(":%v", viper.GetString("port"))
	portString := ":4005"
	fmt.Printf("[-] Listening on %v\n", portString)
	http.ListenAndServe(portString, controllers.App())
}

// logMsg sets a logging message on the middleware.HttpInformer that is set up for each request
// by the logging middleware
func logMsg(r *http.Request, msg string) {
	httpInformant := context.Get(r, middleware.HttpInformantKey)
	hi := httpInformant.(*middleware.HttpInformant)
	hi.SetMessage(msg)
}
