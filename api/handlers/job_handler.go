package handlers

import (
	"encoding/json"
	"jobs-api/api/models"
	"jobs-api/utils"
	"net/http"
	"strconv"

	"gopkg.in/go-playground/validator.v9"
)

func ListJobs(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("user").(models.User)

	jobs := models.ListJobs(user.ID)

	utils.RespondWithJSON(w, http.StatusOK, map[string][]models.Job{"jobs": jobs})
}

func UpdateJob(w http.ResponseWriter, r *http.Request) {
	var job models.UpdateJobInput

	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, _ := r.Context().Value("user").(models.User)

	id, _ := strconv.Atoi(utils.GetParam(r, "id"))

	job.ID = id
	job.UserID = user.ID

	if err := utils.CreateValidation().Struct(job); err != nil {
		var errorMessages map[string]string = utils.GenerateValidationMessages(err.(validator.ValidationErrors))
		utils.RespondValidationError(w, errorMessages)
		return
	}

	if !models.UpdateJob(job.ID, job.UserID, job) {
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"message": "Job not found"})
		return
	}

	jobUpdated := models.GetJob(job.ID, job.UserID)

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"message": "Job updated successfully", "job": jobUpdated})
}

func GetJob(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("user").(models.User)

	id, _ := strconv.Atoi(utils.GetParam(r, "id"))

	job := models.GetJob(id, user.ID)

	if job.ID == 0 {
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"message": "Job not found"})
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"job": job})
}

func DeleteJob(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("user").(models.User)

	id, _ := strconv.Atoi(utils.GetParam(r, "id"))

	if !models.DeleteJob(id, user.ID) {
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"message": "Job not found"})
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"message": "Job deleted successfully", "id": id})
}

func CreateJob(w http.ResponseWriter, r *http.Request) {
	var job models.CreateJobInput

	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, _ := r.Context().Value("user").(models.User)

	job.UserID = user.ID

	if err := utils.CreateValidation().Struct(job); err != nil {
		var errorMessages map[string]string = utils.GenerateValidationMessages(err.(validator.ValidationErrors))
		utils.RespondValidationError(w, errorMessages)
		return
	}

	createdJob := models.CreateJob(job)

	utils.RespondWithJSON(w, http.StatusCreated, map[string]interface{}{"message": "Job created successfully", "job": createdJob})
}
