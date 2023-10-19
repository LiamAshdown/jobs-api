package main

import (
	"fmt"
	"jobs-api/api"
	"jobs-api/config"
	"jobs-api/database"
	"jobs-api/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.Load("config/config.yaml")

	utils.InitialiseLogger()
	database.IntialiseDB()

	router := mux.NewRouter()

	api.InitRoutes(router)

	port := fmt.Sprintf(":%d", config.GetConfig().App.Port)
	utils.GetLogger().Info(fmt.Sprintf("Starting server on port %s", port))

	router.Handle("/", router)
	http.ListenAndServe(port, router)
}
