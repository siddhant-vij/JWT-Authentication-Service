package config

import (
	"log"
	"os"
	"strconv"
	"time"

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

	config.AccessTokenKey = os.Getenv("ACCESS_TOKEN_KEY")
	at_expiry_minutes, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRES_IN"))
	config.AccessTokenExpiresIn = time.Duration(at_expiry_minutes) * time.Minute
	config.AccessTokenMaxAge, _ = strconv.Atoi(os.Getenv("ACCESS_TOKEN_MAX_AGE"))

	config.RefreshTokenKey = os.Getenv("REFRESH_TOKEN_KEY")
	rt_expiry_minutes, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRES_IN"))
	config.RefreshTokenExpiresIn = time.Duration(rt_expiry_minutes) * time.Minute
	config.RefreshTokenMaxAge, _ = strconv.Atoi(os.Getenv("REFRESH_TOKEN_MAX_AGE"))
}
