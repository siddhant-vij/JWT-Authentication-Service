package main

import (
	"fmt"
	// "log"
	// "net/http"

	// "github.com/siddhant-vij/JWT-Authentication-Service/config"
	// "github.com/siddhant-vij/JWT-Authentication-Service/routes"
	"github.com/siddhant-vij/JWT-Authentication-Service/utils"
)

// var apiConfig *routes.ApiConfig = &routes.ApiConfig{}

// func init() {
// 	config.LoadEnv(apiConfig)
// 	config.ConnectDB(apiConfig)
// 	config.ConnectRedis(apiConfig)
// }

func main() {
	pwd := "jwt-auth-service"
	encPwd := utils.EncryptPassword(pwd)
	fmt.Println("Encrypted Password: ", encPwd)
	// Unique on each run: salt (hashing)

	userInputPwd1 := "temp-password"
	fmt.Println(utils.ComparePassword(encPwd, userInputPwd1))
	// false

	userInputPwd2 := "jwt-auth-service"
	fmt.Println(utils.ComparePassword(encPwd, userInputPwd2))
	// true

	userInputPwd3 := "jwt_auth_service"
	fmt.Println(utils.ComparePassword(encPwd, userInputPwd3))
	// false

	userInputPwd4 := "jwt-auth-services"
	fmt.Println(utils.ComparePassword(encPwd, userInputPwd4))
	// false

	// serverAddr := "localhost:" + apiConfig.Port
	// log.Fatal(http.ListenAndServe(serverAddr, nil))
}
