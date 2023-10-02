package user

import (
	"errors"

	"github.com/M-Xue/go-auth-server/server"
)

func UserSignUp(server server.Server, firstName string, lastName string, email string, username string, password string) (*User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}
	newUser, err := createUser(server, firstName, lastName, email, username, hashedPassword)
	return newUser, err
}

func UserLogin(server server.Server, email string, password string) (*User, error) {
	user, err := getUserByEmail(server, email)
	if err != nil {
		return nil, err
	}
	if checkPasswordHash(password, user.Password) {
		return user, nil
	} else {
		return nil, &InvalidCredentialsError{}
	}
}

func GetUserById(server server.Server, id string) (*User, error) {
	user, err := getUserById(server, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UserIsEmailAvailable(server server.Server, email string) (bool, error) {
	_, err := getUserByEmail(server, email)
	if err != nil {
		var userNotFoundError *UserNotFoundError
		if errors.As(err, &userNotFoundError) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func UserIsUsernameAvailable(server server.Server, username string) (bool, error) {
	_, err := getUserByUsername(server, username)
	if err != nil {
		var userNotFoundError *UserNotFoundError
		if errors.As(err, &userNotFoundError) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}
