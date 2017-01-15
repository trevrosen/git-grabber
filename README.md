# git-grabber

## A demo app
Sometimes the best way to get a new language into your organization is with a port of something you've already got.

I've used this stack and structure a few times for microservices. I find it easy to deal with and easy to use to carry forward some of the warts that might exist in legacy projects.

## Stack

* [Gorilla Mux](https://github.com/gorilla/mux) - URL routing
* [Logrus](https://github.com/sirupsen/logrus) - configurable logging
* [Negroni](https://github.com/urfave/negroni) - middleware management
* [Gorm](https://github.com/jinzhu/gorm) - ORM (map SQL tables to structs)
* [Ginkgo](https://github.com/onsi/ginkgo) - BDD-style testing (just for DB unit tests - don't freak)

## Structure
Here's how I structure into subpackages and directories:

* **controllers**
	* `app.go` - routes and the attendant middleware wire-up
	* `github_user_controller.go` - an example of dividing up the `controllers` package into handlers corresponding to groups of routes

* **db**
	* `db.go` - interface and concrete types and connection establishment
	* `github_user.go` - an example of dividing up the `db` package into type-specific files

* **middleware**
	* `request_id.go` - an example of using middleware to tag incoming requests with a unique string
	* `logging.go` - an example of using Logrus-powered middleware as a replacement for `log`

* **_testing**
	* `main_suite.go` - setup for a Ginkgo test suite
	* `github_user_test.go` - an example of using Ginkgo for database unit testing of concrete-type DB methods

