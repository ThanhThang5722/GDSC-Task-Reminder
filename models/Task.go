package models

import (
	"TaskReminder/pkg/database"
	"errors"
	"time"
)

type Task struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

type ReceivedTask struct {
	Task
	Pri string `json:"priority"`
}

func NewTask(title, description string, deadline time.Time) *Task {
	return &Task{
		Title:       title,
		Description: description,
		Deadline:    deadline,
	}
}

func (task *ReceivedTask) InsertTask(userID int) error {
	db := database.GetDbInstance()
	query := `
		INSERT INTO TASK
		VALUES (?,?,?,?,?)
	`
	_, err := db.Exec(query, task.Title, task.Description, task.Deadline, task.Pri, userID)
	if err != nil {
		return errors.New("can't not inserted Task to Database")
	}
	return nil
}
