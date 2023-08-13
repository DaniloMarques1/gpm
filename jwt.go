package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func Generate(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 10)
	claims["authorized"] = true
	claims["user"] = userId

	// TODO move secret somewhere else
	str, err := token.SignedString([]byte("somerandomstring"))
	if err != nil {
		return "", err
	}

	return str, nil
}

func Validate(token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("Something wrong with the signing method\n")
		}

		// TODO move secret somewhere else
		return []byte("somerandomstring"), nil
	})

	if err != nil {
		return "", err
	}

	claims := t.Claims.(jwt.MapClaims)
	userIdIf := claims["userId"]
	userId, ok := userIdIf.(string)
	if !ok {
		return "", fmt.Errorf("Something went wrong getting the user id")
	}

	return userId, nil
}
