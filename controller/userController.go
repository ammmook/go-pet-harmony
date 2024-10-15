package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	managementdb "github.com/f1nn-ach/pj-golang/managementDB"
	"github.com/f1nn-ach/pj-golang/model"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

type Messege map[string]string

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

func UserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		pwd := r.FormValue("password")

		if email == "" || pwd == "" {
			tmpl, err := template.ParseFiles("view/login.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			mess := Messege{
				"message": "Email and password cannot be empty",
			}
			tmpl.Execute(w, mess)
		}

		user, err := managementdb.GetUserByEmail(email)
		if err != nil || user == nil {
			tmpl, err := template.ParseFiles("view/login.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			mess := Messege{
				"message": "User not found",
			}
			tmpl.Execute(w, mess)
		}

		if user.Password != pwd {
			tmpl, err := template.ParseFiles("view/login.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			mess := Messege{
				"message": "Invalid email or password",
			}
			tmpl.Execute(w, mess)
		} else {
			session, err_s := store.Get(r, "session-name")
			if err_s != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			session.Values["authenticated"] = true
			session.Values["user"] = *user

			err_s = session.Save(r, w)
			if err_s != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			tmpl, err := template.ParseFiles("view/result.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			err = tmpl.Execute(w, user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	} else {
		http.ServeFile(w, r, "view/login.html")
	}
}

func GetResultPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "view/result.html")
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
			w.Write(usersJSON)
		}
	}
}
