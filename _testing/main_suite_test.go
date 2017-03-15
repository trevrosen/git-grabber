package main_test

import (
	"fmt"
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
	"github.com/trevrosen/git-grabber/db"

	"testing"
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Main Suite")
}

var (
	sdb *db.SqlDB
	err error
)

var _ = BeforeSuite(func() {
	setConfig()
	sdb, err = db.NewSqlDB()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(sdb)

	sdb.Gorm.AutoMigrate(&db.GitHubUser{})
})

var _ = AfterEach(func() {
	err := truncateTables()
	if err != nil {
		panic(fmt.Errorf("Error resetting DB between tests -- %s", err))
	}
})

// currentCounts returns the count for a db/table combination
func currentCounts(db *db.SqlDB, table string) int {
	var count int
	db.Gorm.Table(table).Count(&count)
	return count
}

// currentCountsAccountables returns the undeleted count for the accountables table
func currentCountsAccountables(db *db.SqlDB) int {
	var count int
	db.Gorm.Table("git_hub_users").Where("deleted_at IS NULL").Count(&count)
	return count
}

// truncateTables truncates the given table in both DBs
func truncateTables() error {
	err := sdb.Gorm.Exec("TRUNCATE TABLE git_hub_users").Error
	if err != nil {
		return err
	}
	return nil
}

// setConfig sets up configuration for the app
func setConfig() error {
	viper.SetDefault("environment", "test")
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.git-grabber")
	return viper.ReadInConfig()
}
