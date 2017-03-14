package controllers

import "testing"

// POST /users
// when the payload is valid, contains correct keys and isn't a dupe, responds with 201
func TestUserCreationValidPayload_NonDupe(t *testing.T) {
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
}

// POST /users
// when the payload can't be turned into a struct responds with 400
func TestUserCreationBadJson(t *testing.T) {
}

// GET /users/{username}
// when the username is in the DB, respond with 200
func TestGetUser_Found(t *testing.T) {
}

// GET /users/{username}
// when the username can't be found responds with 404
func TestGetUser_NotFound(t *testing.T) {
}

// GET /users/{username}
// when the username can be found, the app responds with the proper media type structure
func TestGetUser_MediaTypeStructureIsValid(t *testing.T) {
}
