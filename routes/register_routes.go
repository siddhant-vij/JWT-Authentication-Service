package routes

import (
	"net/http"

	"github.com/siddhant-vij/JWT-Authentication-Service/config"
	"github.com/siddhant-vij/JWT-Authentication-Service/middlewares"
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

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", login)
	mux.Handle("/logout", middlewares.AuthMiddleware(http.HandlerFunc(logout)))
	mux.HandleFunc("/verify", verify)
	mux.Handle("/revoke/{refresh_token}", middlewares.AuthAdmin(http.HandlerFunc(revoke)))
}
