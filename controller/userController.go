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

var store = sessions.NewCookieStore([]byte("fin"))

type Message map[string]string

// Handler for user registration
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

// Handler for user login
func UserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		pwd := r.FormValue("password")

		if email == "" || pwd == "" {
			renderTemplate(w, "login.html", Message{"message": "Email and password cannot be empty"})
			return
		}

		user, err := managementdb.GetUserByEmail(email)
		if err != nil || user == nil {
			renderTemplate(w, "login.html", Message{"message": "User not found"})
			return
		}

		if user.Password != pwd {
			renderTemplate(w, "login.html", Message{"message": "Invalid email or password"})
			return
		}

		// Create session
		session, _ := store.Get(r, "session-name")
		session.Values["user"] = user.Email
		session.Save(r, w)

		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}

	http.ServeFile(w, r, "view/login.html")
}

// Handler for rendering the index page
func IndexPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	userEmail := session.Values["user"]

	// If no session, redirect to login
	if userEmail == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Retrieve user info from the database using email
	user, err := managementdb.GetUserByEmail(userEmail.(string))
	if err != nil || user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Render the index page and pass the user data
	renderTemplate(w, "index.html", user)
}

// Function to handle rendering the result page with session data
func GetResultPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	userEmail := session.Values["user"]

	if userEmail == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Retrieve user information based on session data (email)
	user, err := managementdb.GetUserByEmail(userEmail.(string))
	if err != nil || user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Render the result page with the user information
	renderTemplate(w, "result.html", user)
}

func GetIndexPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	userEmail := session.Values["user"]

	if userEmail == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Retrieve user information based on session data (email)
	user, err := managementdb.GetUserByEmail(userEmail.(string))
	if err != nil || user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Render the result page with the user information
	renderTemplate(w, "index.html", user)
}

// Function to render templates
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("view/" + tmpl)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

// Handler to get all users as JSON
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

// Logout handler
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	delete(session.Values, "user") // Remove user from session
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
