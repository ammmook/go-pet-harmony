package managementdb

import (
	"fmt"

	"github.com/f1nn-ach/pj-golang/initializers"
	"github.com/f1nn-ach/pj-golang/model"
)

func AddBooking(book model.Booking, pet_id int) (int64, error) {
	db := initializers.OpenConnection()
	defer db.Close()

	query := "INSERT INTO Bookings (start_date, end_date, requests, pet_id) VALUES (?,?,?,?)"
	result, err := db.Exec(query,
		book.StartDate,
		book.EndDate,
		book.Request,
		pet_id,
	)
	if err != nil {
		fmt.Println("Database insert error:", err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
