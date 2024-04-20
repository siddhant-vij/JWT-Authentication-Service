package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/siddhant-vij/JWT-Authentication-Service/config"
	"github.com/siddhant-vij/JWT-Authentication-Service/database"
	"github.com/siddhant-vij/JWT-Authentication-Service/routes"
	"github.com/siddhant-vij/JWT-Authentication-Service/utils"
)

var apiConfig *routes.ApiConfig = &routes.ApiConfig{}

func init() {
	config.LoadEnv(apiConfig)
	config.ConnectDB(apiConfig)
}

func insertData() {
	user := database.InsertUserParams{
		ID:           uuid.New(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Email:        "temp@example.com",
		PasswordHash: "password123",
	}
	apiConfig.DBQueries.InsertUser(context.TODO(), user)
}

func getData(w http.ResponseWriter, r *http.Request) {
	insertData()

	credentials := database.GetUserByCredentialsParams{
		Email:        "temp@example.com",
		PasswordHash: "password123",
	}
	user, err := apiConfig.DBQueries.GetUserByCredentials(context.TODO(), credentials)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, user)
}

func main() {
	http.HandleFunc("/api/db/get", getData)

	serverAddr := "localhost:" + apiConfig.Port
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
