package main

import (
	"Task_01/internal/workers"
	"Task_01/pkg/csv_file"
	fileLogger "Task_01/pkg/file_logger"
	"fmt"
	"time"
)

const outFile = "./assets/out.txt"

// Maybe overengineered, but I wanted to make smth like a big project
func main() {
	s := time.Now()
	if data, err := csvFile.ParseFile("./assets/data.csv"); err != nil { //getting data separated in slices from csv file
		panic(err)
	} else {
		if parsedWorkers, err := workers.ParseData(data); err != nil { //getting struct worker that contains all data about workers
			panic(err)
		} else {
			var sortedWorkers []workers.Worker = workers.MapToSortedSlice(parsedWorkers) //sorting workers by data, surname, name
			str := make([]string, len(sortedWorkers))
			for i := 0; i < len(str); i++ {
				str[i] = sortedWorkers[i].String()
			}
			if err = fileLogger.LogSlice(str, outFile); err != nil { //writing into file at the end
				panic(err)
			}
		}
	}
	fmt.Println(time.Since(s))
}
