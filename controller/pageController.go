package controller

import (
	"html/template"
	"log"
	"net/http"

	managementdb "github.com/f1nn-ach/pj-golang/managementDB"
	"github.com/f1nn-ach/pj-golang/model"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("fin"))

type TemplateData struct {
	User     *model.User
	Pets     []model.Pet
	Pet      *model.Pet
	Booking  *model.Booking
	Bookings []model.Booking
	DayCount int
	Message  string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data *TemplateData) {
	t, err := template.ParseFiles("view/"+tmpl, "view/header.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func GetResultPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	userEmail, ok := session.Values["user"].(string)

	if !ok || userEmail == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	user, err := managementdb.GetUserByEmail(userEmail)
	if err != nil || user == nil {
		log.Printf("Error fetching user: %v", err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	log.Printf("User fetched successfully: %+v", user) // Log the entire user struct

	data := &TemplateData{
		User: user,
	}
	renderTemplate(w, "result.html", data)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		pwd := r.FormValue("password")

		if email == "" || pwd == "" {
			renderTemplate(w, "login.html", &TemplateData{
				Message: "Email and password cannot be empty"})
			return
		}

		user, err := managementdb.GetUserByEmail(email)
		if err != nil || user == nil || user.Password != pwd {
			renderTemplate(w, "login.html", &TemplateData{
				Message: "Invalid email or password"})
			return
		}

		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Session retrieval error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		session.Values["user"] = user.Email
		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   900,
			HttpOnly: true,
		}

		if err := session.Save(r, w); err != nil {
			http.Error(w, "Error saving session: "+err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("User logged in successfully: %s", user.Email)
		http.Redirect(w, r, "/result", http.StatusSeeOther)
		return
	}

	renderTemplate(w, "login.html", &TemplateData{Message: ""})
}

// IndexPage handles the index page
func IndexPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	userEmail, ok := session.Values["user"].(string)

	if !ok || userEmail == "" {
		// User not logged in, render index page
		renderTemplate(w, "index.html", &TemplateData{
			Message: ""})
		return
	}

	// Fetch user by email from database
	user, err := managementdb.GetUserByEmail(userEmail)
	if err != nil || user == nil {
		log.Printf("Error fetching user: %v", err)
		renderTemplate(w, "index.html", &TemplateData{
			Message: ""})
		return
	}

	// Render index with user data
	data := &TemplateData{
		User: user,
	}
	renderTemplate(w, "index.html", data)
}

// UserLogout handles user logout
func UserLogout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, "Session retrieval error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Delete user session and save
	delete(session.Values, "user")
	if err := session.Save(r, w); err != nil {
		http.Error(w, "Error saving session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to index page with a logout success message
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
