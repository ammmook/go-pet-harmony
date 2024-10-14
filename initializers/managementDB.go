package initializers

import (
	"database/sql"
	"fmt"

	"github.com/f1nn-ach/pj-golang/model"
)

func GetUsers(email string) (*model.User, error) {
	query := "SELECT * FROM Users WHERE email = ?"

	result := DB.QueryRow(query, email)

	var user model.User
	var role sql.NullString

	err := result.Scan(&user.Email, &user.Firstname, &user.Lastname, &user.PhoneNumber, &user.Password, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with email: %s", email)
		}
		return nil, err
	}

	if role.Valid {
		user.Role = role.String
	} else {
		user.Role = "default"
	}

	return &user, nil
}
