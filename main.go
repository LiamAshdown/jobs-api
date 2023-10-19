package main

import (
	"fmt"
	"jobs-api/api"
	"jobs-api/database"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.IntialiseDB()

	router := mux.NewRouter()

	api.InitRoutes(router)

	port := ":8080"
	fmt.Printf("Server is listening on %s\n", port)
	router.Handle("/", router)
	http.ListenAndServe(port, router)
}
