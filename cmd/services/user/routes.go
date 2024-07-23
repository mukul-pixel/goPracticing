package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"example.com/go-practicing/cmd/auth"
	"example.com/go-practicing/cmd/types"
	"example.com/go-practicing/cmd/utils"

)

type Handler struct {
	Store types.UserStore
}

func NewHandler(Store types.UserStore) *Handler {
	return &Handler{Store: Store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//parsing JSON Payload
	var payload types.RegisterPayLoad
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err) //i would be checking,if my body is null or not and decode the req.body for every req, therefore we'll create reusable functions in utils
	}

	//checking if user already exists
	_, err := h.Store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email:%s already exists", payload.Email))
		return
	}

	//using bcrypt to hash password
	hashedPassword,err:= auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteJSON(w,http.StatusInternalServerError,err)
	}

	//if not exists - create user
	err = h.Store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteError(w,http.StatusCreated,nil)

}
