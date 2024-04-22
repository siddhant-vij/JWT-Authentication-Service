package routes

import (
	"net/http"

	"github.com/siddhant-vij/JWT-Authentication-Service/controllers"
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

func revoke(w http.ResponseWriter, r *http.Request) {
	refreshToken := r.PathValue("refresh_token")
	rtDetails, err := utils.ValidateToken(refreshToken, apiConfig.RefreshTokenKey)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	apiConfig.Tokens[1] = rtDetails
	err = controllers.RevokeRefreshToken(apiConfig)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
}
