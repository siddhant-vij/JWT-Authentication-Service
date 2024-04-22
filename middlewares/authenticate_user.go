package middlewares

import (
	"context"
	"database/sql"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/siddhant-vij/JWT-Authentication-Service/database"
	"github.com/siddhant-vij/JWT-Authentication-Service/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rtCookie, _ := r.Cookie("refresh_token")

		if rtCookie == nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Only logged in users can access this route")
			return
		}

		rtKey := os.Getenv("REFRESH_TOKEN_KEY")
		rtDetails, err := utils.ValidateToken(rtCookie.Value, rtKey)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		userId := rtDetails.UserID

		databaseURL := os.Getenv("DATABASE_URL")
		db, _ := sql.Open("postgres", databaseURL)
		dBQueries := database.New(db)
		uuid := uuid.MustParse(userId)

		_, err = dBQueries.GetUserByID(context.TODO(), uuid)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "User belonging to this token no logger exists")
			return
		}

		next.ServeHTTP(w, r)
	})
}
