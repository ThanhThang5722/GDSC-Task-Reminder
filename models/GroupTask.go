package models

import (
	"TaskReminder/pkg/database"
)

type DoFirst struct {
	ListTask []Task
}

func (group *DoFirst) GetPriority() int {
	return 1
}

type DoLater struct {
	ListTask []Task
}

func (group *DoLater) GetPriority() int {
	return 2
}

type Delegate struct {
	ListTask []Task
}

func (group *Delegate) GetPriority() int {
	return 3
}

type Eliminate struct {
	ListTask []Task
}

func (group *Eliminate) GetPriority() int {
	return 4
}

func GetGroupTask(priority string, userID int) (DoFirst, error) {
	var Res DoFirst
	db := database.GetDbInstance()
	query := `
		SELECT Title, Description, Deadline
		FROM TASK
		WHERE Priority = ?
		AND userID = ?
	`
	resulst, err := db.Query(query, priority, userID)
	if err != nil {
		return Res, err
	}
	defer resulst.Close()
	for resulst.Next() {
		var task Task
		err := resulst.Scan(&task.Title, &task.Description, &task.Deadline)
		if err != nil {
			return Res, err
		}
		Res.ListTask = append(Res.ListTask, task)
	}

	return Res, nil
}
