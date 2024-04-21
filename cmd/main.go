package main

import (
	"log"
	"net/http"

	"github.com/siddhant-vij/JWT-Authentication-Service/config"
	"github.com/siddhant-vij/JWT-Authentication-Service/routes"
)

var apiConfig *routes.ApiConfig = &routes.ApiConfig{}

func init() {
	config.LoadEnv(apiConfig)
	config.ConnectDB(apiConfig)
	config.ConnectRedis(apiConfig)
}

func main() {
	serverAddr := "localhost:" + apiConfig.AuthServerPort
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
