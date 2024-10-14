package managementdb

import (
	"database/sql"

	"github.com/f1nn-ach/pj-golang/initializers"
	"github.com/f1nn-ach/pj-golang/model"
)

func GetUsersByEmail(email string) (*model.User, error) {
	query := "Select * from Users where email = ?"
	result := initializers.DB.QueryRow(query, email)

	var user model.User
	var role sql.NullString

	err := result.Scan(&user.Email, &user.Firstname, &user.Lastname, &user.PhoneNumber, &user.Password, &role)
	if err != nil {
		return nil, err
	}

	if role.Valid {
		user.Role = role.String
	} else {
		user.Role = "default"
	}

	return &user, nil
}
