package controllers

import (
	"context"

	"github.com/siddhant-vij/JWT-Authentication-Service/config"
)

func RevokeRefreshToken(config *config.ApiConfig) error {
	refreshToken := config.Tokens[1]
	err := config.RedisClient.Del(context.TODO(), refreshToken.TokenUuid).Err()
	if err != nil {
		return err
	}
	return nil
}

func IsRTRevoked(config *config.ApiConfig) bool {
	refreshToken := config.Tokens[1]
	_, err := config.RedisClient.Get(context.TODO(), refreshToken.TokenUuid).Result()
	return err != nil
}
