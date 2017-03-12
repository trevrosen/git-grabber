package main_test

import (
	"fmt"

	"github.com/bypasslane/clortho/db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"

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
	sdb = db.NewSqlDB()
})

var _ = AfterEach(func() {
	err := truncateTables()
	if err != nil {
		panic(fmt.Errorf("Error resetting DB between tests -- %s", err))
	}
})

// currentCounts returns the count for a db/table combination
func currentCounts(db *db.Cdb, table string) int {
	var count int
	db.Gorm.Table(table).Count(&count)
	return count
}

// currentCountsAccountables returns the undeleted count for the accountables table
func currentCountsAccountables(db *db.Cdb) int {
	var count int
	db.Gorm.Table("accountables").Where("deleted_at IS NULL").Count(&count)
	return count
}

// truncateTables truncates the given table in both DBs
func truncateTables() error {
	err := apidb.Gorm.Exec("TRUNCATE TABLE git_hub_users").Error
	if err != nil {
		return err
	}
	return nil
}

// setConfig sets up configuration for the app
func setConfig() error {
	viper.SetDefault("environment", "test")
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.clortho")
	return viper.ReadInConfig()
}
