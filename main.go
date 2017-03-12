package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
	"github.com/trevrosen/git-grabber/controllers"
	"github.com/trevrosen/git-grabber/db"
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
