package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	configs "example.com/go-practicing/cmd/config"

)

func CreateJWT(secret []byte, userId int) (string, error) {

	expiration := time.Second * time.Duration(configs.Envs.JWTExpirationTime)

	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"userId" : strconv.Itoa(userId),
		"expiredAt" : time.Now().Add(expiration).Unix(),
	})

	tokenString,err := token.SignedString(secret)
	if err != nil {
		return "" , err
	}

	return tokenString, nil

}