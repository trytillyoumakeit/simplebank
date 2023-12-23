package token

import (
	"testing"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/stretchr/testify/require"
)

func TestPasetoToken(t *testing.T) {
	// Initialize PasetoMaker with a symmetric key
	symmetricKey := paseto.NewV4SymmetricKey() // In real-world, use a secure, constant key
	maker := NewPasetoMaker(symmetricKey, "your-implicit-assertion")

	username := "testuser"
	duration := time.Minute

	// Test successful token creation
	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// Test token verification
	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, username, payload.Username)

	// Test token expiration
	expiredToken, err := maker.CreateToken(username, -time.Minute) // Negative duration for immediate expiration
	require.NoError(t, err)
	require.NotEmpty(t, expiredToken)

	_, err = maker.VerifyToken(expiredToken)
	require.Error(t, err)
	// Optionally check for specific error type if you have custom error handling
}

// Additional test cases can be written for other edge cases
