package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	managementdb "github.com/f1nn-ach/pj-golang/managementDB"
	"github.com/f1nn-ach/pj-golang/model"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		user := model.User{
			Email:       r.FormValue("email"),
			Firstname:   r.FormValue("firstname"),
			Lastname:    r.FormValue("lastname"),
			PhoneNumber: r.FormValue("phoneNumber"),
			Password:    r.FormValue("password"),
		}

		managementdb.AddtUsers(user)

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else if r.Method == http.MethodGet {
		http.ServeFile(w, r, "view/register.html")
	}
}

func CallUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		usersJSON, err := managementdb.GetAllUsers()
		if err != nil {
			log.Fatalf("Error getting users: %v", err)
		}

		var users []model.User
		if err := json.Unmarshal(usersJSON, &users); err != nil {
			log.Fatalf("Error unmarshalling users: %v", err)
		}

		for _, user := range users {
			fmt.Println(user.Firstname)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(usersJSON)
	}
}
