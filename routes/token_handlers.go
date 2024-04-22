package routes

import (
	"net/http"

	"github.com/siddhant-vij/JWT-Authentication-Service/utils"
)

func verify(w http.ResponseWriter, r *http.Request) {
	atCookie, err := r.Cookie("access_token")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	atDetails, err := utils.ValidateToken(atCookie.Value, apiConfig.AccessTokenKey)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, atDetails)
}

// func revoke(w http.ResponseWriter, r *http.Request) {
// }
