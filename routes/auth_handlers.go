package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/siddhant-vij/JWT-Authentication-Service/controllers"
	"github.com/siddhant-vij/JWT-Authentication-Service/utils"
)

func login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	type credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		IsAdmin  bool   `json:"is_admin"`
	}
	user := credentials{}
	err := decoder.Decode(&user)
	if err != nil {
		apiConfig.AuthStatus = "false: "
		utils.RespondWithError(w, http.StatusBadRequest, apiConfig.GetAuthStatus() + err.Error())
		return
	}
	if user.Email == "" || user.Password == "" {
		apiConfig.AuthStatus = "false: "
		utils.RespondWithError(w, http.StatusBadRequest, apiConfig.GetAuthStatus() + "invalid credentials")
		return
	}

	atCookie, err1 := r.Cookie("access_token")
	rtCookie, err2 := r.Cookie("refresh_token")
	if err1 != nil || err2 != nil || atCookie.Value == "" || rtCookie.Value == "" || controllers.IsRTRevoked(apiConfig) {
		err = controllers.RegisterUser(user.Email, user.Password, user.IsAdmin, apiConfig)
		if err != nil {
			apiConfig.AuthStatus = "false: "
			utils.RespondWithError(w, http.StatusBadRequest, apiConfig.GetAuthStatus() + err.Error())
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
			apiConfig.AuthStatus = "false: "
			utils.RespondWithError(w, http.StatusBadRequest, apiConfig.GetAuthStatus() + err.Error())
			return
		}
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    apiConfig.Tokens[0].Token,
		Path:     "/",
		MaxAge:   apiConfig.AccessTokenMaxAge * 60,
		Expires:  time.Now().Add(time.Duration(apiConfig.AccessTokenExpiresIn)),
		Secure:   false,
		HttpOnly: true,
		Domain:   "localhost",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    apiConfig.Tokens[1].Token,
		Path:     "/",
		MaxAge:   apiConfig.RefreshTokenMaxAge * 60,
		Expires:  time.Now().Add(time.Duration(apiConfig.RefreshTokenExpiresIn)),
		Secure:   false,
		HttpOnly: true,
		Domain:   "localhost",
	})

	apiConfig.AuthStatus = "true"
	utils.RespondWithJSON(w, http.StatusOK, apiConfig.GetAuthStatus())
}

func logout(w http.ResponseWriter, r *http.Request) {
	rtCookie, err := r.Cookie("refresh_token")
	if err != nil {
		apiConfig.AuthStatus = "false: "
		utils.RespondWithError(w, http.StatusBadRequest, apiConfig.GetAuthStatus() + err.Error())
		return
	} else {
		rtDetails, _ := utils.ValidateToken(rtCookie.Value, apiConfig.RefreshTokenKey)
		apiConfig.Tokens[1] = rtDetails
		controllers.LogoutUser(apiConfig)

		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    "",
			Path:     "/",
			MaxAge:   0,
			Expires:  time.Unix(0, 0),
			Secure:   false,
			HttpOnly: true,
			Domain:   "localhost",
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    "",
			Path:     "/",
			MaxAge:   0,
			Expires:  time.Unix(0, 0),
			Secure:   false,
			HttpOnly: true,
			Domain:   "localhost",
		})

		apiConfig.AuthStatus = "false: "
		utils.RespondWithJSON(w, http.StatusOK, apiConfig.GetAuthStatus() + "User logged out successfully")
	}
}
