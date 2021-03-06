.PHONY: build check doc fmt lint install run test vet

excluding_vendor := $(shell go list ./... | grep -v /vendor/)


GIT_GRABBER_ENVIRONMENT ?= development

GIT_GRABBER_DB_HOST := $(shell jq --raw-output .$(GIT_GRABBER_ENVIRONMENT).host < ~/.git-grabber/config.json)
GIT_GRABBER_DB_NAME := $(shell jq --raw-output .$(GIT_GRABBER_ENVIRONMENT).name < ~/.git-grabber/config.json)
GIT_GRABBER_DB_PASS := $(shell jq --raw-output .$(GIT_GRABBER_ENVIRONMENT).pass < ~/.git-grabber/config.json)
GIT_GRABBER_DB_PORT := $(shell jq --raw-output .$(GIT_GRABBER_ENVIRONMENT).port < ~/.git-grabber/config.json)
GIT_GRABBER_DB_USER := $(shell jq --raw-output .$(GIT_GRABBER_ENVIRONMENT).user < ~/.git-grabber/config.json)

DEPEND=\
	github.com/onsi/ginkgo \
	github.com/onsi/ginkgo/ginkgo \
	github.com/onsi/gomega \
	github.com/go-sql-driver/mysql \
	github.com/pressly/goose/cmd/goose \
	github.com/spf13/viper \
	github.com/spf13/pflag

default: build

# Build all packages
build:
	go build -i -o git-grabber

check:
	make test && make run

install_deps:
	go get -u $(DEPEND)
	go install $(DEPEND)

run:
	go build && ./git-grabber

db_create:
	mysql -u$(GIT_GRABBER_DB_USER) -p$(GIT_GRABBER_DB_PASS) -e "CREATE DATABASE $(GIT_GRABBER_DB_NAME)"

test:
	go test -v $(excluding_vendor)
	cd _testing && ginkgo

# Start documentation server with search
doc:
	godoc -http=:8080 -index

vet:
	go vet ./...
