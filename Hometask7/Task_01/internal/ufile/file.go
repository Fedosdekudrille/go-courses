package ufile

import (
	"Task_01/pkg/udb"
	"bufio"
	"os"
)

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ParseUsers(lines []string) ([]udb.User, error) {
	var users []udb.User
	var errs []error
	for _, line := range lines {
		user := udb.User{}
		err := user.Parse(line)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		users = append(users, user)
	}
	return users, LogErrors(errs)
}
