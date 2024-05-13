package data

import (
	"errors"
	"strconv"
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	IsCompleted bool   `json:"is_completed"`
}

func (t Task) String() string {
	if t.IsCompleted {
		return strconv.Itoa(t.ID) + ". " + t.Name + " (completed)"
	}
	return strconv.Itoa(t.ID) + ". " + t.Name + " (not completed)"
}

func GetTaskNumById(tasks []Task, id int) (int, error) {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			return i, nil
		}
	}
	return -1, errors.New("no such task")
}

func CompleteTask(tasks []Task, id int) error {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			if tasks[i].IsCompleted {
				return errors.New("task already completed")
			}
			tasks[i].IsCompleted = true
			break
		}
	}
	return nil
}
