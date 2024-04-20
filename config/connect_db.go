package config

import (
	"database/sql"
	"log"
	"os"
	"os/exec"

	_ "github.com/lib/pq"

	"github.com/siddhant-vij/JWT-Authentication-Service/database"
	"github.com/siddhant-vij/JWT-Authentication-Service/routes"
)

func ConnectDB(config *routes.ApiConfig) {
	upCommand := exec.Command("bash", "../scripts/up.sh")
	err := upCommand.Run()
	if err != nil {
		log.Fatal("Error running up.sh: ", err)
		os.Exit(1)
	}

	sqlcCommand := exec.Command("bash", "../scripts/sqlc.sh")
	err = sqlcCommand.Run()
	if err != nil {
		log.Fatal("Error running sqlc.sh: ", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
		os.Exit(1)
	}
	config.DBQueries = database.New(db)
}
