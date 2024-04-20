package main

import (
	"log"
	"net/http"

	"github.com/siddhant-vij/JWT-Authentication-Service/config"
	"github.com/siddhant-vij/JWT-Authentication-Service/routes"
	"github.com/siddhant-vij/JWT-Authentication-Service/utils"
)

func health(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, "OK!")
}

func main() {
	apiConfig := &routes.ApiConfig{}
	config.LoadEnv(apiConfig)

	http.HandleFunc("/api/healthChecker", health)

	serverAddr := "localhost:" + apiConfig.Port
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
