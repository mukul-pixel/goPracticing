package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	configs "example.com/go-practicing/cmd/config"
	"example.com/go-practicing/cmd/types"
	"example.com/go-practicing/cmd/utils"

)

type contextString string
const UserKey contextString = "userId"

func CreateJWT(secret []byte, userId int) (string, error) {

	expiration := time.Second * time.Duration(configs.Envs.JWTExpirationTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(userId),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

//if the user is not login, then there would be no token generated for him
//so will not create his order.

func WithJWTAuth(handlerfunc http.HandlerFunc, userStore types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get token from the user request
		tokenString := getTokenFromRequest(r)
		//validate the token
		token, err := validateJWT(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}
		//fetching the userId from db (id from the token)
		claims := token.Claims.(jwt.MapClaims)
		str := claims["userId"].(string)

		userId, _ := strconv.Atoi(str)
		u, err := userStore.GetUserByID(userId)
		if err != nil {
			log.Printf("failed to get user by id: %v", userId)
			permissionDenied(w)
			return
		}
		//setcontext to the 'userId' as userId in handlefunc
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey,u.ID)
		r=r.WithContext(ctx)

		handlerfunc(w,r)
	}
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}

	return userID
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(configs.Envs.JWTSecret), nil
	})
}
func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorisation")

	if tokenAuth != "" {
		return tokenAuth
	}
	return ""
}
