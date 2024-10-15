package controller

import (
	"math/rand"
	"net/http"
	"time"

	managementdb "github.com/f1nn-ach/pj-golang/managementDB"
	"github.com/f1nn-ach/pj-golang/model"
)

func GenerateID() string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	id := "H"
	for i := 0; i < 5; i++ {
		randomChar := charset[rand.Intn(len(charset))]
		id += string(randomChar)
	}
	return id
}

func PetRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		session, _ := store.Get(r, "session-name")
		userEmail := session.Values["user"]

		pet := model.Pet{
			Name:    r.FormValue("pet_name"),
			Gender:  r.FormValue("pet_gender"),
			Age:     r.FormValue("pet_age"),
			Breed:   r.FormValue("breed"),
			Species: r.FormValue("species"),
		}
		managementdb.AddPets(pet, userEmail.(string))

		http.Redirect(w, r, "/listpets", http.StatusSeeOther)
	} else if r.Method == http.MethodGet {
		http.ServeFile(w, r, "view/registerpet.html")
	}
}
