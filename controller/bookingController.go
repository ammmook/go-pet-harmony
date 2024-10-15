package controller

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"strconv"
	"time"
  
func init() {
	gob.Register(model.Booking{})
}

func BookingRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		pet_id, err := strconv.Atoi(r.FormValue("pet"))
		if err != nil {
			http.Error(w, "Invalid pet id", http.StatusBadRequest)
			return
		}
		startdate := r.FormValue("startDate")
		starttime := r.FormValue("startTime")
		enddate := r.FormValue("endDate")
		endtime := r.FormValue("endTime")

		startDateTimeStr := fmt.Sprintf("%s %s", startdate, starttime)
		endDateTimeStr := fmt.Sprintf("%s %s", enddate, endtime)

		booking := model.Booking{
			StartDate: startDateTimeStr,
			EndDate:   endDateTimeStr,
			Request:   r.FormValue("requests"),
		}
		insertedID, err_regis := managementdb.AddBooking(booking, pet_id)
		booking.Id = int(insertedID)
		if err_regis != nil {
			http.Error(w, err_regis.Error(), http.StatusBadRequest)
			return
		}
		startDateTime, _ := time.Parse("2006-01-02 15:04", startDateTimeStr)
		endDateTime, _ := time.Parse("2006-01-02 15:04", endDateTimeStr)
		duration := endDateTime.Sub(startDateTime)
		dayCount := int(duration.Hours()/24) + 1

		session, _ := store.Get(r, "session-name")
		session.Values["booking"] = booking
		session.Values["dayCount"] = dayCount
		session.Save(r, w)

		http.Redirect(w, r, "/receipt", http.StatusSeeOther)

	} else if r.Method == http.MethodGet {
		session, _ := store.Get(r, "session-name")
		userEmail := session.Values["user"].(string)
		user, err1 := managementdb.GetUserByEmail(userEmail)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}

		pets, err := managementdb.GetPetsByEmail(userEmail)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := &TemplateData{
			User: user,
			Pets: pets,
		}
		renderTemplate(w, "booking.html", data)
	}
}

func BookingDetails(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	userEmail := session.Values["user"].(string)
	user, err1 := managementdb.GetUserByEmail(userEmail)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}

	booking, ok := session.Values["booking"].(model.Booking)
	if !ok {
		http.Error(w, "No booking information found", http.StatusBadRequest)
		return
	}

	dayCount, _ := session.Values["dayCount"].(int)
	data := &TemplateData{
		User:     user,
		Booking:  &booking,
		DayCount: dayCount,
	}

	delete(session.Values, "booking")
	delete(session.Values, "dayCount")
	session.Save(r, w)

	renderTemplate(w, "receipt.html", data)
}
