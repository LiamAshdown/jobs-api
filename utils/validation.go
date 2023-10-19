package utils

import (
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

func CreateValidation() *validator.Validate {
	return validator.New()
}

func GenerateValidationMessages(validationErrors validator.ValidationErrors) map[string]string {
	errorMessages := make(map[string]string)

	for _, err := range validationErrors {
		key := strings.ToLower(err.Field())
		field := err.Field()
		tag := err.Tag()
		param := err.Param()

		switch tag {
		case "required":
			errorMessages[key] = field + " is required"
		case "min":
			errorMessages[key] = field + " must be at least " + param + " characters long"
		case "max":
			errorMessages[key] = field + " must be at most " + param + " characters long"
		case "email":
			errorMessages[key] = field + " must be a valid email address"
		default:
			errorMessages[key] = field + " is invalid [" + tag + "]"
		}
	}

	return errorMessages
}
