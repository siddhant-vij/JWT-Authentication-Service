package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/siddhant-vij/JWT-Authentication-Service/routes"
)

func LoadEnv(config *routes.ApiConfig) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.PostgresHost = os.Getenv("POSTGRES_HOST")
	config.PostgresPort = os.Getenv("POSTGRES_PORT")
	config.PostgresUser = os.Getenv("POSTGRES_USER")
	config.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	config.PostgresDB = os.Getenv("POSTGRES_DB")

	config.DatabaseURL = os.Getenv("DATABASE_URL")

	config.Port = os.Getenv("PORT")
	config.ResourceOrigin = os.Getenv("RESOURCE_ORIGIN")
	config.ClientOrigin = os.Getenv("CLIENT_ORIGIN")

	config.RedisUrl = os.Getenv("REDIS_URL")

	config.AccessTokenPrivateKey = os.Getenv("ACCESS_TOKEN_PRIVATE_KEY")
	config.AccessTokenPublicKey = os.Getenv("ACCESS_TOKEN_PUBLIC_KEY")
	config.AccessTokenExpiredIn = os.Getenv("ACCESS_TOKEN_EXPIRED_IN")
	config.AccessTokenMaxAge, _ = strconv.Atoi(os.Getenv("ACCESS_TOKEN_MAX_AGE"))

	config.RefreshTokenPrivateKey = os.Getenv("REFRESH_TOKEN_PRIVATE_KEY")
	config.RefreshTokenPublicKey = os.Getenv("REFRESH_TOKEN_PUBLIC_KEY")
	config.RefreshTokenExpiredIn = os.Getenv("REFRESH_TOKEN_EXPIRED_IN")
	config.RefreshTokenMaxAge, _ = strconv.Atoi(os.Getenv("REFRESH_TOKEN_MAX_AGE"))
}
