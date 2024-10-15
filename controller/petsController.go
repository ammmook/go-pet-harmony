package controller

import (
	"net/http"
	"strconv"

	managementdb "github.com/f1nn-ach/pj-golang/managementDB"
	"github.com/f1nn-ach/pj-golang/model"
)

func LoadListPetPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		session, _ := store.Get(r, "session-name")
		userEmail, _ := session.Values["user"].(string)

		pets, err := managementdb.GetPetsByEmail(userEmail)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := &TemplateData{
			Pets: pets,
		}
		renderTemplate(w, "listpet.html", data)
	}
}

func PetRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		session, _ := store.Get(r, "session-name")
		userEmail := session.Values["user"].(string)

		pet := model.Pet{
			Name:    r.FormValue("pet_name"),
			Gender:  r.FormValue("pet_gender"),
			Age:     r.FormValue("pet_age"),
			Breed:   r.FormValue("breed"),
			Species: r.FormValue("species"),
		}
		managementdb.AddPets(pet, userEmail)
		http.Redirect(w, r, "/listpets", http.StatusSeeOther)
	} else if r.Method == http.MethodGet {
		renderTemplate(w, "registerpet.html", &TemplateData{})
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

		data := &TemplateData{
			Pet: pet,
		}

		renderTemplate(w, "editpet.html", data)
	}
}

func DeletePet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "Invalid pet id", http.StatusBadRequest)
		return
	}
	managementdb.DeletePet(id)
	http.Redirect(w, r, "/listpets", http.StatusSeeOther)
}
