package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"example.com/go-practicing/cmd/types"

)

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserHandler{}
	userHandler := NewHandler(userStore)

	t.Run(
		"our main should fail if we give invalid payload and this should pass",
		func(t *testing.T) {
		payload := types.RegisterPayLoad{
			FirstName: "Mukul",
			LastName:  "Khatri",
			Email:     "invalid",
			Password:  "helloworld",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", userHandler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run(
		"should correctly register the user",
		func(t *testing.T) {
		payload := types.RegisterPayLoad{
			FirstName: "Mukul",
			LastName:  "Khatri",
			Email:     "asd@gmail.com",
			Password:  "helloworld",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", userHandler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockUserHandler struct{}

// CreateUser implements types.UserStore.
func (m *mockUserHandler) CreateUser(types.User) error {
	return nil
}

// GetUSerById implements types.UserStore.
func (m *mockUserHandler) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

// GetUserByEmail implements types.UserStore.
func (m *mockUserHandler) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found with email:%s", email)
}
