package controller

import (
	"log"

	"github.com/f1nn-ach/pj-golang/initializers"
	"github.com/f1nn-ach/pj-golang/model"
)

func GetUser(userID int) (*model.User, error) {
	query := `SELECT email, firstname, lastname, phone_number, password, role FROM users WHERE id = ?`
	row := initializers.DB.QueryRow(query, userID)

	var user model.User

	err := row.Scan(&user.Email, &user.Firstname, &user.Lastname, &user.PhoneNumber, &user.Password, &user.Role)
	if err != nil {
		log.Println("Error scanning user:", err)
		return nil, err
	}

	return &user, nil
}
