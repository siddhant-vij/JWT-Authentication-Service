package routes

import (
	"net/http"

	"github.com/siddhant-vij/JWT-Authentication-Service/config"
)

var apiConfig *config.ApiConfig = &config.ApiConfig{}

func init() {
	config.LoadEnv(apiConfig)
	config.ConnectDB(apiConfig)
	config.ConnectRedis(apiConfig)
}

func AuthServerPort() string {
	return apiConfig.AuthServerPort
}

func RegisterRoutes() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
}
