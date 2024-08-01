package auth

import (
	"fmt"
	"testing"

	configs "example.com/go-practicing/cmd/config"

)

func TestJwtCreate(t *testing.T) {
	secret := []byte(configs.Envs.JWTSecret)

	token, err := CreateJWT(secret, 1)
	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}

	if token == "" {
		t.Error("expected token to be not empty")
	}

	fmt.Printf("token:%v\n", token)
}
