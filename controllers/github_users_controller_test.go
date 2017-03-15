package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/trevrosen/git-grabber/db"
)

// POST /users
// when the payload is valid, contains correct keys and isn't a dupe, responds with 201
func TestUserCreationValidPayload_NonDupe(t *testing.T) {
	sdb := &db.MockSdb{
		OnFindGitHubUserByName: func(name string) (*db.GitHubUser, error) { return nil, nil },
		OnSaveRecord:           func(record interface{}) error { return nil },
	}

	server := httptest.NewServer(App(sdb))
	defer server.Close()
	payload := `{"username":"joebob"}`
	res, err := postCreateUser(server, payload)

	if err != nil {
		log.Fatalln(err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 but received %d", res.StatusCode)
	}
}

// POST /users
// when the payload is valid, the media type returned is of the proper structure
func TestUserCreationValidPayload_MediaTypeStructureIsValid(t *testing.T) {
}

// POST /users
// when the payload is valid, contains correct keys and represents a duplicate, responds with 409
func TestUserCreationValidPayload_Dupe(t *testing.T) {
}

// POST /users
// when the payload doesn't contain the right keys responds with 400
func TestUserCreationInValidPayload(t *testing.T) {
	sdb := &db.MockSdb{
		OnFindGitHubUserByName: func(name string) (*db.GitHubUser, error) { return nil, nil },
		OnSaveRecord:           func(record interface{}) error { return nil },
	}

	server := httptest.NewServer(App(sdb))
	defer server.Close()
	payload := `{"name":"joebob"}`
	res, err := postCreateUser(server, payload)

	if err != nil {
		log.Fatalln(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected 400 but received %d", res.StatusCode)
	}
}

// POST /users
// when the payload can't be turned into a struct responds with 400
func TestUserCreationBadJson(t *testing.T) {
	sdb := &db.MockSdb{
		OnFindGitHubUserByName: func(name string) (*db.GitHubUser, error) { return nil, nil },
		OnSaveRecord:           func(record interface{}) error { return nil },
	}

	server := httptest.NewServer(App(sdb))
	defer server.Close()
	payload := `"username":"joebob"}`
	res, err := postCreateUser(server, payload)

	if err != nil {
		log.Fatalln(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected 400 but received %d", res.StatusCode)
	}
}

// GET /users/{username}
// when the username is in the DB, respond with 200
func TestGetUser_Found(t *testing.T) {
	sdb := &db.MockSdb{
		OnFindGitHubUserByName: validUser,
	}

	server := httptest.NewServer(App(sdb))
	defer server.Close()
	res, err := getUser(server, "joebob")

	if err != nil {
		log.Fatalln(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 but received %d", res.StatusCode)
	}
}

// GET /users/{username}
// when the username can't be found responds with 404
func TestGetUser_NotFound(t *testing.T) {
	sdb := &db.MockSdb{
		OnFindGitHubUserByName: func(name string) (*db.GitHubUser, error) { return nil, gorm.ErrRecordNotFound },
	}

	server := httptest.NewServer(App(sdb))
	defer server.Close()
	res, err := getUser(server, "joebob")

	if err != nil {
		log.Fatalln(err)
	}

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Expected 404 but received %d", res.StatusCode)
	}
}

// GET /users/{username}
// when the username can be found, the app responds with the proper media type structure
func TestGetUser_MediaTypeStructureIsValid(t *testing.T) {
}

// postCreateUser sends an HTTP POST to create a GitHubUser
func postCreateUser(server *httptest.Server, githubUserJson string) (*http.Response, error) {
	client := &http.Client{}
	payload := strings.NewReader(githubUserJson)
	req, _ := http.NewRequest("POST", server.URL+"/users", payload)
	return client.Do(req)
}

// getUser sends an HTTP GET to fetch a GitHubUser
func getUser(server *httptest.Server, username string) (*http.Response, error) {
	client := &http.Client{}
	return client.Get(server.URL + "/users/" + username)
}
func validUser(name string) (*db.GitHubUser, error) {
	return &db.GitHubUser{
		Username: name,
	}, nil
}
