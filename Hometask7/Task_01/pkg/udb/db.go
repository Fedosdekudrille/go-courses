package udb

import (
	"database/sql"
	_ "modernc.org/sqlite"
	"os"
)

const (
	dbPath           = "./db/users.db"
	createScriptPath = "./db/scripts/db_create.sql"
)

func EnsureCreatedDB() error {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	data, err := os.ReadFile(createScriptPath)
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = db.Exec(string(data))
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func ClearUsers() error {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("DELETE FROM users")
	return err
}

func GetUsers() ([]User, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, email, password_hash, username, is_active FROM users")
	if err != nil {
		return nil, err
	}
	users := []User{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Email, &user.PasswordHash, &user.Username, &user.IsActive)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func AddUsers(users []User) error {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	for _, user := range users {
		_, err = db.Exec("INSERT INTO users (email, password_hash, username, is_active) VALUES (?, ?, ?, ?)", user.Email, user.PasswordHash, user.Username, user.IsActive)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

//the next are unused, additional

func GetUser(id int) (User, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return User{}, err
	}
	defer db.Close()
	row := db.QueryRow("SELECT id, email, password_hash, username, is_active FROM users WHERE id = ?", id)
	user := User{}
	err = row.Scan(&user.Id, &user.Email, &user.PasswordHash, &user.Username, &user.IsActive)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func AddUser(user User) error {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO users (email, password_hash, username, is_active) VALUES (?, ?, ?, ?)", user.Email, user.PasswordHash, user.Username, user.IsActive)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func UpdateUser(user User) error {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE users SET email = ?, password_hash = ?, username = ?, is_active = ? WHERE id = ?", user.Email, user.PasswordHash, user.Username, user.IsActive, user.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func DeleteUser(id int) error {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
