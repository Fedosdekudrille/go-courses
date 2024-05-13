package workers

import (
	fileLogger "Task_01/pkg/file_logger"
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	errorFile = "assets/errors.txt"
	layout    = "1/2/2006"
)

func ParseData(data [][]string) (res map[string]*Worker, err error) {
	res = make(map[string]*Worker, len(data))
	var strErr []string
	for id, str := range data {
		if name, err := checkName(str[0]); err == nil {
			if startTime, err := time.Parse(layout, str[1]); err == nil {
				if endTime, err := time.Parse(layout, str[2]); err == nil {
					if wVac, ok := res[str[0]]; ok {
						wVac.AddVacationDate(startTime, endTime)
					} else {
						res[str[0]] = NewWorker(name[0], name[1], [][2]time.Time{{startTime, endTime}})
					}
					continue // if everything is successful, we don't log to the errors.txt
				}
			}
		}
		strErr = append(strErr, strconv.Itoa(id+1)+" "+str[0]+","+str[1]+","+str[2])
	}
	if err = fileLogger.LogSlice(strErr, errorFile); err != nil {
		return nil, err
	}
	return res, nil
}

func checkName(name string) ([]string, error) {
	str := strings.Split(name, " ")
	if len(str) != 2 { // check whether name contains of two parts: name and surname
		return nil, errors.New("wrong worker name")
	}
	return str, nil
}
