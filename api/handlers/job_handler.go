package handlers

import (
	"encoding/json"
	"net/http"
)

func ListJobs(w http.ResponseWriter, r *http.Request) {
	jobs, _ := json.Marshal([]string{"job1", "job2", "job3"})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(jobs)
}
