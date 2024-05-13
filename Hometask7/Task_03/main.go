package main

import (
	"Task_02/internal/data"
	"Task_02/internal/parm"
	"Task_02/internal/server"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	mode, val := parm.GetFlags()
	switch mode {
	case parm.READ:
		tasks, err := data.GetTasksJson("assets/tasks.json")
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("No tasks found")
			return
		} else if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, task := range tasks {
			fmt.Println(task)
		}
	case parm.COMPLETE:
		tasks, err := data.GetTasksJson("assets/tasks.json")
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("No tasks found")
			return
		} else if err != nil {
			fmt.Println(err.Error())
			return
		}
		id, err := strconv.Atoi(val)

		err = data.CompleteTask(tasks, id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = data.SaveTasksJson("assets/tasks.json", tasks)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	case parm.WRITE:
		tasks, err := data.GetTasksJson("assets/tasks.json")
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			fmt.Println(err.Error())
			return
		}
		if len(tasks) == 0 {
			tasks = append(tasks, data.Task{ID: 1, Name: val, IsCompleted: false})
		} else {
			tasks = append(tasks, data.Task{ID: tasks[len(tasks)-1].ID + 1, Name: val, IsCompleted: false})
		}
		err = data.SaveTasksJson("assets/tasks.json", tasks)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	case parm.SERVER:
		server.MustStartServer()
	default:
		fmt.Println("Specify mode: -server, -list, -task [task name], -complete [task Id]")
	}
}
