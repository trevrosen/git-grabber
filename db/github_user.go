package db

import "github.com/jinzhu/gorm"

// GitHubUser represents a user of GitHub (shockingly)
type GitHubUser struct {
	gorm.Model
	Username string
}

// FindGitHubUserByName returns the DB record of the provided GitHub user
func (sdb *SqlDB) FindGitHubUserByName(name string) (*GitHubUser, error) {
	user := &GitHubUser{}
	err := sdb.Gorm.Where("username = ?", name).First(user).Error
	return user, err
}
