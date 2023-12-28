package auth

import (
	"time"

	"github.com/google/uuid"
)

// Interface for managing authentication tokens
type TokenFactory interface {
	CreateAuthToken(userID uuid.UUID, duration time.Duration) (string, error)
	VerifyAndParseAuthToken(token string) (AuthTokenPayload, error)
}
