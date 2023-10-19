package api

import (
	"jobs-api/api/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	jobsRouter := router.PathPrefix("/jobs").Subrouter()
	jobsRouter.HandleFunc("", handlers.ListJobs).Methods("GET")

	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
}
