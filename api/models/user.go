package models

import (
	"jobs-api/database"
)

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserCredentials struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func DoesUserAlreadyExist(email string) bool {
	db := database.GetDB()

	var exists bool
	db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", email).Scan(&exists)

	return exists
}

func GetUser(userId int) User {
	db := database.GetDB()

	user := User{}

	db.QueryRow("SELECT id, email, created_at, updated_at FROM users WHERE id = ?", userId).Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	return user
}

func CreateUser(user UserCredentials) User {
	db := database.GetDB()

	result, err := db.Exec("INSERT INTO users (email, password, created_at, updated_at) VALUES (?, ?, NOW(), NOW())", user.Email, user.Password)

	if err != nil {
		panic(err.Error())
	}

	id, _ := result.LastInsertId()

	return GetUser(int(id))
}

func GetUserByEmail(email string) User {
	db := database.GetDB()

	user := User{}

	db.QueryRow("SELECT id, email, password, created_at, updated_at FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	return user
}
