package managementdb

import (
	"database/sql"
	"fmt"

	"github.com/f1nn-ach/pj-golang/initializers"
	"github.com/f1nn-ach/pj-golang/model"
)

func AddBooking(book model.Booking, pet_id int) (sql.Result, error) {
	db := initializers.OpenConnection()

	query := "INSERT INTO Bookings (start_date, end_date, request, pet_id) VALUES (?,?,?,?)"
	result, err := db.Exec(query,
		book.StartDate,
		book.EndDate,
		book.Request,
		pet_id,
	)
	if err != nil {
		fmt.Println("Database insert error:", err)
		return nil, err
	}

	return result, nil
}
