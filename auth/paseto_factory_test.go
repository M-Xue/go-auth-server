package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPasetoFactory(t *testing.T) {
	factory, err := NewPasetoFactory("12345678901234567890123456789012")
	require.NoError(t, err)

	userID, err := uuid.NewUUID()
	require.NoError(t, err)
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := factory.CreateAuthToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := factory.VerifyAndParseAuthToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.TokenID)
	require.Equal(t, userID, payload.UserID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}
