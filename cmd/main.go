package main

import (
	"log"
	"net/http"

	"github.com/siddhant-vij/JWT-Authentication-Service/routes"
)

func main() {
	routes.RegisterRoutes()
	
	serverAddr := "localhost:" + routes.AuthServerPort()
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
