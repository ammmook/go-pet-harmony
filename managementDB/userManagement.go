package managementdb

import (
	"database/sql"
	"encoding/json"

	"github.com/f1nn-ach/pj-golang/initializers"
	"github.com/f1nn-ach/pj-golang/model"
)

func GetAllUsers() ([]byte, error) {
	db := initializers.OpenConnection()
	defer db.Close()

	var users []model.User
	result, err := db.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		var user model.User
		var role sql.NullString
		if err := result.Scan(&user.Email, &user.Firstname, &user.Lastname, &user.Password, &user.PhoneNumber, &role); err != nil {
			return nil, err
		}

		if role.Valid {
			user.Role = role.String
		} else {
			user.Role = "default"
		}

		users = append(users, user)
	}

	res, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func AddtUsers(user model.User) (sql.Result, error) {
	db := initializers.OpenConnection()
	defer db.Close()

	query := "INSERT INTO Users (email, firstname, lastname, phone_number, password) VALUES (?, ?, ?, ?, ?)"
	result, err := db.Exec(query,
		user.Email,
		user.Firstname,
		user.Lastname,
		user.PhoneNumber,
		user.Password,
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetUserByEmail(email string) (*model.User, error) {
	db := initializers.OpenConnection()
	defer db.Close()

	query := "SELECT * FROM Users WHERE email = ?"

	result := db.QueryRow(query, email)

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
