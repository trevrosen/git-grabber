package db

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// DBInteractor is the interface for talking to a database
type DBInteractor interface {
	FindGitHubUserByName(string) GitHubUser
}

// SqlDB implements DBInteractor for a Gorm
type SqlDB struct {
	// Gorm is the accessor to do Gorm stuff
	Gorm *gorm.DB
	// ConnectionString returns a string in the form of "mysql://<USER>:<PASS>@<HOST>/<DB_NAME>"
	ConnectionString string
}

// NewSqlDB initializes a SqlDB with a connection to the SQL DB
func NewSqlDB(dbName string) (*SqlDB, error) {
	ndb := &SqlDB{
		ConnectionString: dbString(dbName),
	}
	return ndb, ndb.establishConnection()
}

// establishConnection sets up the Gorm connection to the SQL DB
func (sdb *SqlDB) establishConnection() error {
	var err error
	sdb.Gorm, err = gorm.Open("mysql", sdb.ConnectionString)
	return err
}

func dbString(dbName string) string {
	environment := viper.GetString("environment")
	user := viper.GetString(strings.Join([]string{dbName, environment, "user"}, "."))
	pass := viper.GetString(strings.Join([]string{dbName, environment, "pass"}, "."))
	host := viper.GetString(strings.Join([]string{dbName, environment, "host"}, "."))
	port := viper.GetString(strings.Join([]string{dbName, environment, "port"}, "."))
	name := viper.GetString(strings.Join([]string{dbName, environment, "name"}, "."))
	return fmt.Sprintf("%v:%v@(%v:%v)/%v?parseTime=true", user, pass, host, port, name)
}
