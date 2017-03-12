package controllers

import "testing"

// POST /users
// when the payload is valid, contains correct keys and isn't a dupe, responds with 201
func TestUserCreationValidPayload_NonDupe(t *testing.T) {
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
