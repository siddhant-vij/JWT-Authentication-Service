package routes

import (
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/siddhant-vij/JWT-Authentication-Service/database"
)

type ApiConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string

	DatabaseURL string

	Port           string
	ResourceOrigin string
	ClientOrigin   string

	RedisUrl string

	AccessTokenKey       string
	AccessTokenExpiresIn time.Duration
	AccessTokenMaxAge    int

	RefreshTokenKey       string
	RefreshTokenExpiresIn time.Duration
	RefreshTokenMaxAge    int

	DBQueries   *database.Queries
	RedisClient *redis.Client
}
