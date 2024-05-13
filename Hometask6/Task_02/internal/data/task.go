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

func CompleteTask(tasks []Task, id int) error {
	if id < 1 || id > len(tasks) {
		return errors.New("no such task")
	}
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
