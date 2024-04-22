package controllers

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/siddhant-vij/JWT-Authentication-Service/config"
	"github.com/siddhant-vij/JWT-Authentication-Service/database"
	"github.com/siddhant-vij/JWT-Authentication-Service/utils"
)

func RegisterUser(email, password string, isAdmin bool, config *config.ApiConfig) error {
	user, errUser := config.DBQueries.GetUserByEmail(context.TODO(), email)

	if errUser == nil && !utils.ComparePassword(user.PasswordHash, password) {
		err := config.RedisClient.Del(context.TODO(), config.Tokens[1].TokenUuid).Err()
		if err != nil {
			return err
		}
		return errors.New("invalid password")
	}

	var userId uuid.UUID
	newUserId := uuid.New()
	if user.ID == uuid.Nil {
		userId = newUserId
	} else {
		userId = user.ID
	}
	curTime := time.Now()

	if errUser != nil {
		var insertUserParams = database.InsertUserParams{
			ID:           userId,
			CreatedAt:    curTime,
			UpdatedAt:    curTime,
			Email:        email,
			PasswordHash: utils.EncryptPassword(password),
			IsAdmin:      isAdmin,
		}
		_, err := config.DBQueries.InsertUser(context.TODO(), insertUserParams)
		if err != nil {
			return err
		}
	}

	atDetails, err := utils.CreateToken(userId.String(), config.AccessTokenExpiresIn, config.AccessTokenKey)
	if err != nil {
		return err
	}
	rtDetails, err := utils.CreateToken(userId.String(), config.RefreshTokenExpiresIn, config.RefreshTokenKey)
	if err != nil {
		return err
	}

	if config.Tokens[1].TokenUuid != "" {
		err := config.RedisClient.Del(context.TODO(), config.Tokens[1].TokenUuid).Err()
		if err != nil {
			return err
		}
	}

	config.Tokens[0] = atDetails
	config.Tokens[1] = rtDetails

	err = config.RedisClient.Set(context.TODO(), rtDetails.TokenUuid, rtDetails.UserID, time.Duration(rtDetails.ExpiresIn)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

var userIdAt string
var userIdRt string

func LoginUser(config *config.ApiConfig, errList []error) error {
	atDetails := config.Tokens[0]
	errAt := errList[0]
	rtDetails := config.Tokens[1]
	errRt := errList[1]

	if errAt == nil {
		config.Tokens[0] = atDetails
		userIdAt = atDetails.UserID
	} else {
		atDetailsNew, err := utils.CreateToken(userIdAt, config.AccessTokenExpiresIn, config.AccessTokenKey)
		if err != nil {
			return err
		}
		config.Tokens[0] = atDetailsNew
	}

	if errRt == nil {
		config.Tokens[1] = rtDetails
		userIdRt = rtDetails.UserID
		err := config.RedisClient.Set(context.TODO(), rtDetails.TokenUuid, rtDetails.UserID, time.Duration(rtDetails.ExpiresIn)*time.Second).Err()
		if err != nil {
			return err
		}
	} else {
		rtDetailsNew, err := utils.CreateToken(userIdRt, config.RefreshTokenExpiresIn, config.RefreshTokenKey)
		if err != nil {
			return err
		}
		config.Tokens[1] = rtDetailsNew

		if !strings.Contains(errRt.Error(), "token has expired") {
			err = config.RedisClient.Del(context.TODO(), rtDetails.TokenUuid).Err()
			if err != nil {
				return err
			}
		}

		err = config.RedisClient.Set(context.TODO(), rtDetailsNew.TokenUuid, rtDetailsNew.UserID, time.Duration(rtDetailsNew.ExpiresIn)*time.Second).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func LogoutUser(config *config.ApiConfig) error {
	rtDetails := config.Tokens[1]
	err := config.RedisClient.Del(context.TODO(), rtDetails.TokenUuid).Err()
	if err != nil {
		return err
	}
	return nil
}
