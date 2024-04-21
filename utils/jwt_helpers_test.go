package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTokenWithCorrectFormatAndStructure(t *testing.T) {
	user_id := "123"
	ttl := time.Hour
	privateKey := "secret"

	tokenDetails, err := CreateToken(user_id, ttl, privateKey)

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenDetails.Token)
	assert.NotEmpty(t, tokenDetails.TokenUuid)
	assert.Equal(t, user_id, tokenDetails.UserID)
	assert.Equal(t, int64(ttl.Seconds()), tokenDetails.ExpiresIn)
}

func TestCreateTokenWithEmptyPrivateKey(t *testing.T) {
	user_id := "123"
	ttl := time.Hour
	privateKey := ""

	tokenDetails, err := CreateToken(user_id, ttl, privateKey)

	assert.Error(t, err)
	assert.Equal(t, TokenDetails{}, tokenDetails)
}

func TestCreateTokenWithEmptyUserID(t *testing.T) {
	user_id := ""
	ttl := time.Hour
	privateKey := "secret"

	tokenDetails, err := CreateToken(user_id, ttl, privateKey)

	assert.Error(t, err)
	assert.Equal(t, TokenDetails{}, tokenDetails)
}

func TestCreateTokenWithZeroTTL(t *testing.T) {
	user_id := "123"
	ttl := 0
	privateKey := "secret"

	tokenDetails, err := CreateToken(user_id, time.Duration(ttl), privateKey)

	assert.Error(t, err)
	assert.Equal(t, TokenDetails{}, tokenDetails)
}

func TestCreateTokenWithNegativeTTL(t *testing.T) {
	user_id := "123"
	ttl := -5
	privateKey := "secret"

	tokenDetails, err := CreateToken(user_id, time.Duration(ttl), privateKey)

	assert.Error(t, err)
	assert.Equal(t, TokenDetails{}, tokenDetails)
}

func TestValidateTokenValidToken(t *testing.T) {
	expected, _ := CreateToken("user_id", 3600, "public_key")
	tokenStr := expected.Token
	publicKey := "public_key"

	result, err := ValidateToken(tokenStr, publicKey)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestValidateTokenEmptyTokenString(t *testing.T) {
	tokenStr := ""
	publicKey := "public_key"

	expected := TokenDetails{}

	result, err := ValidateToken(tokenStr, publicKey)

	assert.Error(t, err)
	assert.Equal(t, expected, result)
}

func TestValidateTokenEmptyKey(t *testing.T) {
	tokenStr := "123"
	publicKey := ""

	expected := TokenDetails{}

	result, err := ValidateToken(tokenStr, publicKey)

	assert.Error(t, err)
	assert.Equal(t, expected, result)
}

func TestValidateTokenInvalidToken(t *testing.T) {
	tokenStr := "invalid_token"
	publicKey := "public_key"

	expected := TokenDetails{}

	result, err := ValidateToken(tokenStr, publicKey)

	assert.Error(t, err)
	assert.Equal(t, expected, result)
}

func TestValidateTokenExpiredToken(t *testing.T) {
  expected, _ := CreateToken("user_id", 1, "public_key")
	tokenStr := expected.Token
  publicKey := "public_key"

	time.Sleep(2 * time.Second) // For token to expire.

  _, err := ValidateToken(tokenStr, publicKey)

  assert.Error(t, err)
}
