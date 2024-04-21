package utils

import (
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	if password == "" {
		panic("password cannot be empty")
	}
	if len([]byte(password)) > 72 {
		panic("password cannot be greater than 72 bytes")
	}
	data, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func encryptPasswordWithCost(password string, cost int) string {
	if password == "" {
		panic("password cannot be empty")
	}
	if len([]byte(password)) > 72 {
		panic("password cannot be greater than 72 bytes")
	}
	data, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func ComparePassword(hashedPassword, password string) bool {
	if hashedPassword == "" || password == "" {
		return false
	}
	if len([]byte(password)) > 72 {
		return false
	}
	if !utf8.ValidString(password) {
		return false
	}
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
