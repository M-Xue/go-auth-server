package token

import "time"

type Factory interface {
	// Creates a new token for a specific username and duration
	CreateAuthToken(username string, duration time.Duration) (string, error)

	// Checks if the token is valid or not
	VerifyAuthToken(token string) (*Payload, error)
}
