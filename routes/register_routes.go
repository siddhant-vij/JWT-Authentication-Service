package routes

import (
	"net/http"

	"github.com/siddhant-vij/JWT-Authentication-Service/config"
	"github.com/siddhant-vij/JWT-Authentication-Service/utils"
)

var apiConfig *config.ApiConfig = &config.ApiConfig{}

func init() {
	config.LoadEnv(apiConfig)
	config.ConnectDB(apiConfig)
	config.ConnectRedis(apiConfig)

	apiConfig.Tokens = make([]utils.TokenDetails, 2)
}

func AuthServerPort() string {
	return apiConfig.AuthServerPort
}

func RegisterRoutes() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
}
