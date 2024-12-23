package models

import (
	"TaskReminder/pkg/database"
	"errors"
	"fmt"
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
		INSERT INTO TASK (Title, Description, Deadline, Priority, userID)
		VALUES (?,?,?,?,?)
	`
	_, err := db.Exec(query, task.Title, task.Description, task.Deadline, task.Pri, userID)
	if err != nil {
		fmt.Println(err)
		return errors.New("can't not insert Task to Database")
	}
	return nil
}

func (task *ReceivedTask) DeleteTask(userID int) error {
	db := database.GetDbInstance()
	query := `
		DELETE FROM TASK
		WHERE userID = ? AND Title = ? AND Description = ?
	`
	_, err := db.Exec(query, userID, task.Title, task.Description)
	if err != nil {
		return errors.New("can't not delete Task to Database")
	}
	return nil
}

func (task *ReceivedTask) UpdateTaskPriority(userID int) error {
	db := database.GetDbInstance()
	query := `
		UPDATE TASK
		SET Priority = ?
		WHERE userID = ? AND Title = ? AND Description = ? AND Deadline = ?
	`
	_, err := db.Exec(query, task.Pri, userID, task.Title, task.Description, task.Deadline)
	if err != nil {
		return errors.New("can't not update Task's priority to Database")
	}
	return nil
}
