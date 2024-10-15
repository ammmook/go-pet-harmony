package controller

import (
	"net/http"
	"strconv"

	managementdb "github.com/f1nn-ach/pj-golang/managementDB"
	"github.com/f1nn-ach/pj-golang/model"
)

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

func LoadListPetPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		session, _ := store.Get(r, "session-name")
		userEmail, _ := session.Values["user"].(string)

		pets, err := managementdb.GetPetsByEmail(userEmail)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		renderTemplate(w, "listpet.html", pets)
	}
}

func EditPet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id, err := strconv.Atoi(r.FormValue("Id"))
		if err != nil {
			http.Error(w, "Invalid pet id", http.StatusBadRequest)
			return
		}
		pet := model.Pet{
			Id:      id,
			Name:    r.FormValue("pet_name"),
			Gender:  r.FormValue("pet_gender"),
			Age:     r.FormValue("pet_age"),
			Breed:   r.FormValue("breed"),
			Species: r.FormValue("species"),
		}
		managementdb.EditPets(pet)

		http.Redirect(w, r, "/listpets", http.StatusSeeOther)
	} else if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Invalid pet id", http.StatusBadRequest)
			return
		}
		pet, err := managementdb.GetPetById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		renderTemplate(w, "editpet.html", pet)
	}
}
