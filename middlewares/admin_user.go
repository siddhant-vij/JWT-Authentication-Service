package middlewares

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/siddhant-vij/JWT-Authentication-Service/database"
	"github.com/siddhant-vij/JWT-Authentication-Service/utils"
)

func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		type credentials struct {
			Email    string `json:"email"`
			Password string `json:"password"`
			IsAdmin  bool   `json:"is_admin"`
		}
		user := credentials{}
		err := decoder.Decode(&user)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		if user.Email == "" || user.Password == "" {
			utils.RespondWithError(w, http.StatusBadRequest, "Please provide email and password")
			return
		}
		if !user.IsAdmin {
			utils.RespondWithError(w, http.StatusUnauthorized, "Only admins can access this route")
			return
		}

		databaseURL := os.Getenv("DATABASE_URL")
		db, _ := sql.Open("postgres", databaseURL)
		dBQueries := database.New(db)

		resUser, err := dBQueries.GetUserByEmail(context.TODO(), user.Email)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "User belonging to this admin token no logger exists")
			return
		}
		if !resUser.IsAdmin {
			utils.RespondWithError(w, http.StatusUnauthorized, "Only admins can access this route")
			return
		}

		next.ServeHTTP(w, r)
	})
}
