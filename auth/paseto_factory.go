package auth

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type PasetoFactory struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func (factory PasetoFactory) CreateAuthToken(
	userID uuid.UUID,
	duration time.Duration,
) (string, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", err
	}

	return factory.paseto.Encrypt(factory.symmetricKey, payload, nil)
}

func (factory PasetoFactory) VerifyAndParseAuthToken(token string) (AuthTokenPayload, error) {
	payload := AuthTokenPayload{}
	err := factory.paseto.Decrypt(token, factory.symmetricKey, payload, nil)
	if err != nil {
		return AuthTokenPayload{}, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return AuthTokenPayload{}, err
	}

	return payload, nil
}

func NewPasetoFactory(symmetricKey string) (TokenFactory, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf(
			"invalid key size: must be exactly %d characters",
			chacha20poly1305.KeySize,
		)
	}

	pasetoTokenFactory := PasetoFactory{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return pasetoTokenFactory, nil
}
