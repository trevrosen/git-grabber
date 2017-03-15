package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/trevrosen/git-grabber/controllers"
	"github.com/trevrosen/git-grabber/db"
	"github.com/trevrosen/git-grabber/middleware"
)

var sdb *db.SqlDB

func main() {
	var err error
	setConfig()

	// TODO: allow override of environment via flag

	sdb, err = db.NewSqlDB()

	if err != nil {
		log.Println("there was an error establishing DB connection: ", err.Error())
	}
	sdb.Gorm.AutoMigrate(&db.GitHubUser{})
	bindAndRun(sdb)
}

// bindAndRun starts the server
func bindAndRun(sdb *db.SqlDB) {
	//portString := fmt.Sprintf(":%v", viper.GetString("port"))
	portString := ":4005"
	fmt.Printf("[-] Listening on %v\n", portString)
	http.ListenAndServe(portString, controllers.App(sdb))
}

// setConfig sets up configuration for the app
func setConfig() error {
	viper.SetDefault("environment", "development")
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.git-grabber")
	return viper.ReadInConfig()
}

// setupLogging configures a global logger and injects it into packages that support logger injection
func setupLogging() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	// Log everything and colorize logs in development mode
	// You can customize to any level but "info" in development mode
	if viper.GetString("environment") == "development" {
		logger.Formatter = &logrus.TextFormatter{}
		logger.Level = logrus.DebugLevel
	}

	middleware.SetLogger(logger)
}
