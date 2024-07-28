package auth

import (
	"fmt"
	"testing"
)

func TestJwtCreate(t *testing.T) {
	secret := []byte("secret")
	token, err := CreateJWT(secret, 3)
	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}

	if token == "" {
		t.Errorf("didn't expected that token will be empty")
	}

	fmt.Printf("token:%v\n", token)
}
