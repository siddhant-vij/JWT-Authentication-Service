package routes

import (
	"encoding/json"
	"net/http"

	"github.com/siddhant-vij/JWT-Authentication-Service/controllers"
	"github.com/siddhant-vij/JWT-Authentication-Service/utils"
)

func login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	type credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
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

	atCookie, err1 := r.Cookie("access_token")
	rtCookie, err2 := r.Cookie("refresh_token")
	if err1 != nil || err2 != nil {
		err = controllers.RegisterUser(user.Email, user.Password, apiConfig)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		atDetails, errAt := utils.ValidateToken(atCookie.Value, apiConfig.AccessTokenKey)
		apiConfig.Tokens[0] = atDetails

		rtDetails, errRt := utils.ValidateToken(rtCookie.Value, apiConfig.RefreshTokenKey)
		apiConfig.Tokens[1] = rtDetails

		var errList []error = []error{errAt, errRt}
		err = controllers.LoginUser(apiConfig, errList)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    apiConfig.Tokens[0].Token,
		Path:     "/",
		MaxAge:   apiConfig.AccessTokenMaxAge * 60,
		Secure:   false,
		HttpOnly: true,
		Domain:   "localhost",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    apiConfig.Tokens[1].Token,
		Path:     "/",
		MaxAge:   apiConfig.RefreshTokenMaxAge * 60,
		Secure:   false,
		HttpOnly: true,
		Domain:   "localhost",
	})
}

func logout(w http.ResponseWriter, r *http.Request) {
	atCookie, err1 := r.Cookie("access_token")
	rtCookie, err2 := r.Cookie("refresh_token")
	if err1 != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err1.Error())
		return
	} else if err2 != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err2.Error())
		return
	} else {
		atDetails, errAt := utils.ValidateToken(atCookie.Value, apiConfig.AccessTokenKey)
		apiConfig.Tokens[0] = atDetails

		rtDetails, errRt := utils.ValidateToken(rtCookie.Value, apiConfig.RefreshTokenKey)
		apiConfig.Tokens[1] = rtDetails

		var errList []error = []error{errAt, errRt}
		err := controllers.LogoutUser(apiConfig, errList)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
	}
}
