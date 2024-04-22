package config

import (
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/siddhant-vij/JWT-Authentication-Service/database"
	"github.com/siddhant-vij/JWT-Authentication-Service/utils"
)

type ApiConfig struct {
	DatabaseURL    string
	RedisUrl       string
	AuthServerPort string

	AccessTokenKey       string
	AccessTokenExpiresIn time.Duration
	AccessTokenMaxAge    int

	RefreshTokenKey       string
	RefreshTokenExpiresIn time.Duration
	RefreshTokenMaxAge    int

	DBQueries   *database.Queries
	RedisClient *redis.Client

	Tokens  []utils.TokenDetails
}
