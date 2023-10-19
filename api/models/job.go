package models

import (
	"jobs-api/database"
	"jobs-api/date_formatter"
)

type Job struct {
	ID                     int     `json:"id"`
	Title                  string  `json:"title"`
	Description            string  `json:"description"`
	Salary                 float32 `json:"salary"`
	Company                string  `json:"company"`
	UserId                 int     `json:"user_id"`
	CreatedAt              string  `json:"created_at"`
	UpdatedAt              string  `json:"updated_at"`
	HumanReadableCreatedAt string  `json:"human_readable_created_at"`
}

type Jobs []Job

type CreateJobInput struct {
	Title       string `json:"title" validate:"required,min=5,max=255"`
	Description string `json:"description" validate:"required,min=5,max=255"`
	Salary      string `json:"salary" validate:"required"`
	Company     string `json:"company" validate:"required,min=5,max=255"`
	UserID      int    `json:"user_id" validate:"required"`
}

type UpdateJobInput struct {
	CreateJobInput
	ID int `json:"id" validate:"required"`
}

func ListJobs(userId int) Jobs {
	db := database.GetDB()
	rows, _ := db.Query("SELECT id, title, description, salary, company, created_at, updated_at FROM jobs WHERE user_id = ?", userId)

	defer rows.Close()

	jobs := Jobs{}

	for rows.Next() {
		var job Job

		rows.Scan(&job.ID, &job.Title, &job.Description, &job.Salary, &job.Company, &job.CreatedAt, &job.UpdatedAt)

		jobCreatedAt := date_formatter.FormatDate(job.CreatedAt)
		job.HumanReadableCreatedAt = jobCreatedAt

		jobs = append(jobs, job)
	}
	return jobs
}

func GetJob(id int, userId int) Job {
	db := database.GetDB()

	var job Job

	db.QueryRow("SELECT id, title, description, salary, company, created_at, updated_at FROM jobs WHERE id = ? AND user_id = ?", id, userId).Scan(&job.ID, &job.Title, &job.Description, &job.Salary, &job.Company, &job.CreatedAt, &job.UpdatedAt)

	jobCreatedAt := date_formatter.FormatDate(job.CreatedAt)
	job.HumanReadableCreatedAt = jobCreatedAt

	return job
}

func DeleteJob(id int, userId int) bool {
	db := database.GetDB()

	result, _ := db.Exec("DELETE FROM jobs WHERE id = ? AND user_id = ?", id, userId)

	rowsAffected, _ := result.RowsAffected()

	return rowsAffected > 0
}

func UpdateJob(id int, userId int, job UpdateJobInput) bool {
	db := database.GetDB()

	result, _ := db.Exec("UPDATE jobs SET title = ?, description = ?, salary = ?, company = ?, updated_at = NOW() WHERE id = ? AND user_id = ?", job.Title, job.Description, job.Salary, job.Company, id, userId)

	rowsAffected, _ := result.RowsAffected()

	return rowsAffected > 0
}

func CreateJob(job CreateJobInput) Job {
	db := database.GetDB()

	result, _ := db.Exec("INSERT INTO jobs (title, description, salary, company, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, NOW(), NOW())", job.Title, job.Description, job.Salary, job.Company, job.UserID)

	id, _ := result.LastInsertId()

	return GetJob(int(id), job.UserID)
}
