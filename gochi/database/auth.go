package database

import (
	"errors"
	"gohairdresser/structs"

	_ "github.com/go-sql-driver/mysql"
)

var ErrAccountNotFound = errors.New("account not found")
var ErrInvalidPassword = errors.New("invalid password")

func LoginClient(email, password string) (structs.Client, error) {
	acc, err := GetClientByEmail(email)
	if err != nil {
		return structs.Client{}, ErrAccountNotFound
	}

	passValid := CheckPasswordHash(password, acc.Password)
	if !passValid {
		return structs.Client{}, ErrInvalidPassword
	}
	return acc, nil
}
