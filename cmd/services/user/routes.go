package user

import (
	"net/http"

	"github.com/gorilla/mux"

	"example.com/go-practicing/cmd/types"
	"example.com/go-practicing/cmd/utils"

)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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
	if err := utils.ParseJSON(r,&payload); err != nil {
		utils.WriteError(w,http.StatusBadRequest,err); //i would be checking,if my body is null or not and decode the req.body for every req, therefore we'll create reusable functions in utils
	}

	//checking if user already exists
	

}
