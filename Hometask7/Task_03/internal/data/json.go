package data

import (
	"encoding/json"
	"os"
)

func GetTasksJson(path string) ([]Task, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	var tasks []Task
	err = json.NewDecoder(file).Decode(&tasks)
	return tasks, err
}

func SaveTasksJson(path string, tasks []Task) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		return err
	}
	err = json.NewEncoder(file).Encode(tasks)
	return err
}
