package handlers

import (
	"encoding/json"
	"jobs-api/api/models"
	"jobs-api/utils"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserCredentials

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := utils.CreateValidation().Struct(user); err != nil {
		var errorMessages map[string]string = utils.GenerateValidationMessages(err.(validator.ValidationErrors))
		utils.RespondValidationError(w, errorMessages)
		return
	}

	if models.DoesUserAlreadyExist(user.Email) {
		utils.RespondWithJSON(w, http.StatusConflict, map[string]string{"message": "User already exists"})
		return
	}

	// Hash the password
	user.Password, err = utils.HashPassword(user.Password)

	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Something went wrong"})
		return
	}

	createdUser := models.CreateUser(user)

	// Generate a JWT token
	token, err := utils.GenerateJWTToken(createdUser)

	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Something went wrong"})
		return
	}

	// Delete the password from the response
	createdUser.Password = ""

	// All good, return the token
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"message": "User registered successfully", "token": token, "user": createdUser})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserCredentials

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := utils.CreateValidation().Struct(user); err != nil {
		var errorMessages map[string]string = utils.GenerateValidationMessages(err.(validator.ValidationErrors))
		utils.RespondValidationError(w, errorMessages)
		return
	}

	// Check if user exists
	storedUser := models.GetUserByEmail(user.Email)

	if storedUser.Email == "" {
		utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"message": "Invalid email or password"})
		return
	}

	// Compare the passwords
	if !utils.ComparePasswords(storedUser.Password, user.Password) {
		utils.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
		return
	}

	// Generate a JWT token
	token, err := utils.GenerateJWTToken(storedUser)

	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Something went wrong"})
		return
	}

	// Delete the password from the response
	storedUser.Password = ""

	// All good, return the token
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"message": "User logged in successfully", "token": token, "user": storedUser})
}
