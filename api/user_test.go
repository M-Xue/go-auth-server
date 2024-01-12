package api

import (
	"testing"
)

func TestSignUp(t *testing.T) {
	firstName := ""
	lastName := ""
	email := ""
	username := ""
	password := ""

	SignUp(testServer, firstName, lastName, email, username, password)
}

func TestSignUpExistingEmailError(t *testing.T) {
	firstName := ""
	lastName := ""
	email := ""
	username := ""
	password := ""

	SignUp(testServer, firstName, lastName, email, username, password)
}

func TestSignUpExistingUsernameError(t *testing.T) {
	firstName := ""
	lastName := ""
	email := ""
	username := ""
	password := ""

	SignUp(testServer, firstName, lastName, email, username, password)
}
