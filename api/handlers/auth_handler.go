package handlers

import (
	"encoding/json"
	"jobs-api/database"
	"jobs-api/utils"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

type UserCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

var validate = validator.New()

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user UserCredentials

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := validate.Struct(user); err != nil {
		var errorMessages map[string]string = utils.GenerateValidationMessages(err.(validator.ValidationErrors))
		utils.RespondValidationError(w, errorMessages)
		return
	}

	db := database.GetDB()

	// Check to see if user already exists
	var exists bool
	db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", user.Email).Scan(&exists)

	if exists {
		utils.RespondWithJSON(w, http.StatusConflict, map[string]string{"message": "User already exists"})
		return
	}

	// Hash the password
	user.Password, err = utils.HashPassword(user.Password)

	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Something went wrong"})
		return
	}

	// Insert the user into the database
	_, err = db.Exec("INSERT INTO users (email, password, created_at, updated_at) VALUES (?, ?, NOW(), NOW())", user.Email, user.Password)

	// Check for errors inserting the user
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Something went wrong"})
		return
	}

	// Generate a JWT token
	token, err := utils.GenerateJWTToken(user.Email)

	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Something went wrong"})
		return
	}

	// All good, return the token
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"message": "User registered successfully", "token": token, "user": map[string]string{"email": user.Email}})
}
