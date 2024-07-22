package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"example.com/go-practicing/cmd/services/user"

)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

//to run the server
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("listening to:", s.addr)

	return http.ListenAndServe(s.addr, router)
}
