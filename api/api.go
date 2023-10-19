package api

import (
	"jobs-api/api/handlers"
	"jobs-api/api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	// Prefix with api
	apiRouter := router.PathPrefix("/api").Subrouter()

	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	authRouter.Handle("/register", middleware.LoggingMiddleware(http.HandlerFunc(handlers.RegisterUser))).Methods("POST")
	authRouter.Handle("/login", middleware.LoggingMiddleware(http.HandlerFunc(handlers.LoginUser))).Methods("POST")

	jobsRouter := apiRouter.PathPrefix("/jobs").Subrouter()
	jobsRouter.Handle("", middleware.AuthenticationMiddleware(middleware.LoggingMiddleware(http.HandlerFunc(handlers.ListJobs)))).Methods("GET")
	jobsRouter.Handle("", middleware.AuthenticationMiddleware(middleware.LoggingMiddleware(http.HandlerFunc(handlers.CreateJob)))).Methods("POST")
	jobsRouter.Handle("/{id:[0-9]+}", middleware.AuthenticationMiddleware(middleware.LoggingMiddleware(http.HandlerFunc(handlers.UpdateJob)))).Methods("PUT")
	jobsRouter.Handle("/{id:[0-9]+}", middleware.AuthenticationMiddleware(middleware.LoggingMiddleware(http.HandlerFunc(handlers.GetJob)))).Methods("GET")
	jobsRouter.Handle("/{id:[0-9]+}", middleware.AuthenticationMiddleware(middleware.LoggingMiddleware(http.HandlerFunc(handlers.DeleteJob)))).Methods("DELETE")
}
