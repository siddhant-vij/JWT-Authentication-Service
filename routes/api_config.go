package routes

import (
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

	AccessTokenPrivateKey string
	AccessTokenPublicKey  string
	AccessTokenExpiredIn  string
	AccessTokenMaxAge     int

	RefreshTokenPrivateKey string
	RefreshTokenPublicKey  string
	RefreshTokenExpiredIn  string
	RefreshTokenMaxAge     int

	DBQueries *database.Queries
	RedisClient *redis.Client
}
