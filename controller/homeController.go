package controller

import (
	"github.com/f1nn-ach/pj-golang/initializers"
	"github.com/f1nn-ach/pj-golang/model"
)

func GetUsers(userId int) (*model.User, error) {
	query := "SELECT * FROM users WHERE id=?"
	result := initializers.DB.QueryRow(query, userId) // Use QueryRow for single result

	var user model.User
	err := result.Scan(&user.Email, &user.Firstname, &user.Lastname, &user.PhoneNumber, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
