package utils

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetParam(r *http.Request, key string) string {
	params := mux.Vars(r)
	return params[key]
}
