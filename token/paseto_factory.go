package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoFactory struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func (factory *PasetoFactory) CreateAuthToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	return factory.paseto.Encrypt(factory.symmetricKey, payload, nil)
}

func (factory *PasetoFactory) VerifyAuthToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := factory.paseto.Decrypt(token, factory.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func NewPasetoFactory(symmetricKey string) (Factory, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoFactory{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}
