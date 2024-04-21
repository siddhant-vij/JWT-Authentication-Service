package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenDetails struct {
	Token     string
	TokenUuid string
	UserID    string
	ExpiresIn int64
}

func CreateToken(user_id string, ttl time.Duration, key string) (TokenDetails, error) {
	if user_id == "" || key == "" {
		return TokenDetails{}, fmt.Errorf("user id or private key cannot be empty")
	}
	if ttl <= 0 {
		return TokenDetails{}, fmt.Errorf("ttl must be greater than 0")
	}
	tokenUuid := uuid.New().String()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    "jwt-auth-feedagg",
		"IssuedAt":  jwt.NewNumericDate(time.Now()),
		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(ttl)),
		"Subject":   user_id,
		"TokenUuid": tokenUuid,
	})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return TokenDetails{}, err
	}

	return TokenDetails{
		Token:     tokenString,
		TokenUuid: tokenUuid,
		UserID:    user_id,
		ExpiresIn: int64(ttl.Seconds()),
	}, nil
}

func ValidateToken(tokenStr string, key string) (TokenDetails, error) {
	if tokenStr == "" || key == "" {
		return TokenDetails{}, fmt.Errorf("token string or public key cannot be empty")
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		expiresIn := int64(token.Claims.(jwt.MapClaims)["ExpiresAt"].(float64) - float64(time.Now().Unix()))
		if expiresIn < 0 {
			return nil, fmt.Errorf("token has expired")
		}
		
		return []byte(key), nil
	})

	if err != nil {
		return TokenDetails{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return TokenDetails{
			Token:     tokenStr,
			TokenUuid: claims["TokenUuid"].(string),
			UserID:    claims["Subject"].(string),
			ExpiresIn: int64(claims["ExpiresAt"].(float64) - float64(time.Now().Unix())),
		}, nil
	}

	return TokenDetails{}, fmt.Errorf("invalid token")
}
