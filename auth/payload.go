package auth

import (
	"time"

	"github.com/google/uuid"
)

type AuthTokenPayload struct {
	TokenID   uuid.UUID `json:"token_id"`
	UserID    uuid.UUID `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (payload AuthTokenPayload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func NewPayload(userID uuid.UUID, duration time.Duration) (AuthTokenPayload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return AuthTokenPayload{}, err
	}

	payload := AuthTokenPayload{
		TokenID:   tokenID,
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}
