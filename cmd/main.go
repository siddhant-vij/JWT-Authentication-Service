package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/siddhant-vij/JWT-Authentication-Service/config"
	"github.com/siddhant-vij/JWT-Authentication-Service/routes"
	"github.com/siddhant-vij/JWT-Authentication-Service/utils"
)

var apiConfig *routes.ApiConfig = &routes.ApiConfig{}

func init() {
	config.LoadEnv(apiConfig)
	config.ConnectDB(apiConfig)
	config.ConnectRedis(apiConfig)
}

func signupTest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	user := parameters{}
	err := decoder.Decode(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	atDetails, _ := utils.CreateToken(user.Email, apiConfig.AccessTokenExpiresIn, apiConfig.AccessTokenKey)

	rtDetails, _ := utils.CreateToken(user.Email, apiConfig.RefreshTokenExpiresIn, apiConfig.RefreshTokenKey)

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    atDetails.Token,
		Path:     "/",
		MaxAge:   apiConfig.AccessTokenMaxAge * 60,
		Secure:   false,
		HttpOnly: true,
		Domain:   "localhost",
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    rtDetails.Token,
		Path:     "/",
		MaxAge:   apiConfig.RefreshTokenMaxAge * 60,
		Secure:   false,
		HttpOnly: true,
		Domain:   "localhost",
	})

	var response struct {
		ATDetails utils.TokenDetails
		RTDetails utils.TokenDetails
	}
	response.ATDetails = atDetails
	response.RTDetails = rtDetails

	utils.RespondWithJSON(w, http.StatusOK, response)
}

func loginTest(w http.ResponseWriter, r *http.Request) {
	access_token, err := r.Cookie("access_token")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	refresh_token, err := r.Cookie("refresh_token")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	accessTokenDetails, err := utils.ValidateToken(access_token.Value, apiConfig.AccessTokenKey)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	refreshTokenDetails, err := utils.ValidateToken(refresh_token.Value, apiConfig.RefreshTokenKey)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var response struct {
		ATEmail string `json:"at_email"`
		RTEmail string `json:"rt_email"`
	}
	response.ATEmail = accessTokenDetails.UserID
	response.RTEmail = refreshTokenDetails.UserID

	utils.RespondWithJSON(w, http.StatusOK, response)
}

func main() {
	http.HandleFunc("/signup", signupTest)
	http.HandleFunc("/login", loginTest)

	serverAddr := "localhost:" + apiConfig.Port
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
