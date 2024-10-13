package controller

import (
	"database/sql"
	"fmt"

	"github.com/f1nn-ach/pj-golang/initializers"
	"github.com/f1nn-ach/pj-golang/model"
)

func GetUsers(email string) (*model.User, error) {
	query := "SELECT * FROM Users WHERE email = ?"

	result := initializers.DB.QueryRow(query, email)

	var user model.User

	err := result.Scan(&user.Email, &user.Firstname, &user.Lastname, &user.PhoneNumber, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with email: %s", email)
		}
		return nil, err
	}

	return &user, nil
}
