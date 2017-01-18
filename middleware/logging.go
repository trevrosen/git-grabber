package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/sirupsen/logrus"
	"github.com/twinj/uuid"
	"github.com/urfave/negroni"
)

// logger holds the instance of a Logrus Logger
var logger = logrus.New()

// SetLogger sets the logger for this package, allowing dependency injection
func SetLogger(globalLogger *logrus.Logger) {
	logger = globalLogger
}

// GGLogger is a middleware handler that logs the request as it goes in and the response as it goes out.
type GGLogger struct {
	// Logger inherits from log.Logger used to log messages with the Logger middleware
	Logger *logrus.Logger
}

// NewGGLogger returns a new Logger instance
func NewGGLogger() *GGLogger {
	cl := &GGLogger{Logger: logger}
	return cl
}

func (cl *GGLogger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	httpInformant := &HttpInformant{
		id:         uuid.NewV4().String(),
		method:     r.Method,
		path:       r.URL.Path,
		remoteAddr: r.RemoteAddr,
	}

	context.Set(r, HttpInformantKey, httpInformant)

	startTime := time.Now().UTC().Round(time.Millisecond)
	cl.Logger.WithFields(logrus.Fields{
		"id":         httpInformant.ID(),
		"remoteAddr": httpInformant.RemoteAddr(),
		"method":     httpInformant.Method(),
		"path":       httpInformant.Path(),
	}).Info()

	next(rw, r)

	res := rw.(negroni.ResponseWriter)
	endTime := time.Since(startTime)
	cl.Logger.WithFields(logrus.Fields{
		"id":         httpInformant.ID(),
		"statusCode": fmt.Sprintf("%d", res.Status()),
		"statusText": http.StatusText(res.Status()),
		"path":       httpInformant.Path(),
		"duration":   endTime,
	}).Info(httpInformant.Message())
}
