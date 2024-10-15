package controller

// import (
// 	"net/http"

// 	managementdb "github.com/f1nn-ach/pj-golang/managementDB"
// 	"github.com/f1nn-ach/pj-golang/model"
// )

// func BookingRegister(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		session, _ := store.Get(r, "session-name")
// 		userEmail := session.Values["user"].(string)

// 		pets, err := managementdb.GetPetsByEmail(userEmail)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		booking := model.Booking{
// 			StartDate: r.FormValue("startDate"),
// 		}

// 	} else if r.Method == http.MethodGet {
// 		session, _ := store.Get(r, "session-name")
// 		userEmail := session.Values["user"].(string)

// 		pets, err := managementdb.GetPetsByEmail(userEmail)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		data := &TemplateData{
// 			Pets: pets,
// 		}
// 		renderTemplate(w, "booking.html", data)
// 	}
// }
