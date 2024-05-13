package udb

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

type User struct {
	Id           int
	Email        string
	PasswordHash string
	Username     string
	IsActive     bool
}

func (user *User) Parse(line string) error {
	data := strings.Split(line, " ")
	if len(data) != 2 {
		return errors.New("line \"" + line + "\" doesn't match format <name surname>")
	}
	user.Username = line
	user.Email = data[0] + data[1] + "@coolcompany.com"
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Email), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(pass)
	user.IsActive = true
	return nil
}

func (user *User) String() string {
	return "Id: " + strconv.Itoa(user.Id) + "; Username: " + user.Username + "; Email: " + user.Email +
		"; Password: " + user.PasswordHash + "; IsActive: " + strconv.FormatBool(user.IsActive)
}
