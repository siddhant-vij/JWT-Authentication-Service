package main

import (
	"context"
	"log"
	"net/http"

	"github.com/siddhant-vij/JWT-Authentication-Service/config"
	"github.com/siddhant-vij/JWT-Authentication-Service/routes"
	"github.com/siddhant-vij/JWT-Authentication-Service/utils"
)

var apiConfig *routes.ApiConfig = &routes.ApiConfig{}

func init() {
	config.LoadEnv(apiConfig)
	config.ConnectDB(apiConfig)
	config.ConnectRedis(apiConfig)
}

func insertData() {
	key := "key"
	value := "value"

	apiConfig.RedisClient.Set(context.TODO(), key, value, 0)
}

func getData(w http.ResponseWriter, r *http.Request) {
	insertData()

	value, err := apiConfig.RedisClient.Get(context.TODO(), "key").Result()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, value)
}

func main() {
	http.HandleFunc("/api/redis/get", getData)

	serverAddr := "localhost:" + apiConfig.Port
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
