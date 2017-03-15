package main_test

import (
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/trevrosen/git-grabber/db"
)

var _ = Describe("GitHubUser DB methods", func() {
	Describe("FindByUserName", func() {
		var ghu *db.GitHubUser
		var username string = "joebob"

		Describe("when they are in the database", func() {
			BeforeEach(func() {
				ghu = &db.GitHubUser{
					Username: "joebob",
				}
				sdb.SaveRecord(ghu)
			})

			It("should find the user", func() {
				fetchedUser, _ := sdb.FindGitHubUserByName(username)
				Expect(fetchedUser.Username).To(Equal(username))
			})
		})

		Describe("when they are NOT in the database", func() {
			It("should return nil for the user", func() {
				fetchedUser, _ := sdb.FindGitHubUserByName(username)
				Expect(fetchedUser).To(BeNil())
			})

			It("should return gorm.ErrRecordNotFound for the error", func() {
				fetchedUser, err := sdb.FindGitHubUserByName(username)
				Expect(fetchedUser).To(BeNil())
				Expect(err).To(Equal(gorm.ErrRecordNotFound))
			})
		})
	})
})
