package models

import (
	"errors"

	"github.com/kchernenko/eventssample/db"
	"github.com/kchernenko/eventssample/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId

	return err
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var foundPassword string
	err := row.Scan(&user.ID, &foundPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(foundPassword, user.Password)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
