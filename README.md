# git-grabber

## A demo app
Sometimes the best way to get a new language into your organization is with a port of something you've already got.

I've used this stack and structure a few times for microservices. I find it easy to deal with and flexible enough to carry forward some of the (sadly necessary) warts that might exist in legacy projects.

## What it does
* Boot
* Not a whole lot, but not nothin
* Settle down, this is is just here to demo a **structure**

## Prerequisites/Conventions

* Use a Makefile to make common tasks easy
* Use Make to install non-app deps like Ginkgo, Goose, Glide
* Use `jq` for grabbing config info in Make tasks
* Config lives in `$HOME/.<APP_NAME>`
* DB info lives in `config.json` in the config dir
* Glide is the package manager

## Stack

* [Gorilla Mux](https://github.com/gorilla/mux) - URL routing/multiplexing
* [Logrus](https://github.com/sirupsen/logrus) - configurable logging
* [Negroni](https://github.com/urfave/negroni) - middleware management
* [Viper](https://github.com/spf13/viper) - unified configuration
* [Gorm](https://github.com/jinzhu/gorm) - ORM-style DB access
* [Ginkgo](https://github.com/onsi/ginkgo) - BDD-style testing (just for DB unit tests)
* [Glide](https://github.com/Masterminds/glide) - dependency management

## Structure
Here's how I structure into subpackages and directories:

* **controllers**
	* `app.go` - routes and the attendant middleware wire-up
	* `github_user_controller.go` - an example of dividing up the `controllers` package into handlers corresponding to groups of routes

* **db**
	* `db.go` - interface and concrete type declarations and connection establishment
	* `github_user.go` - an example of dividing up the `db` package into type-specific files

* **middleware**
	* `request_id.go` - an example of using middleware to tag incoming requests with a unique string
	* `logging.go` - an example of using Logrus-powered middleware as a replacement for `log`

* **_testing**
	* `main_suite.go` - setup for a Ginkgo test suite
	* `github_user_test.go` - an example of using Ginkgo for database unit testing of concrete-type DB methods

