package main

import (
	"log"
	"net/http"

	"github.com/siddhant-vij/JWT-Authentication-Service/middlewares"
	"github.com/siddhant-vij/JWT-Authentication-Service/routes"
)

func main() {
	mux := http.NewServeMux()
	corsMux := middlewares.CorsMiddleware(mux)
	routes.RegisterRoutes(mux)

	serverAddr := "localhost:" + routes.AuthServerPort()
	log.Fatal(http.ListenAndServe(serverAddr, corsMux))
}
