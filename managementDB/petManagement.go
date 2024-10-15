package managementdb

import (
	"database/sql"

	"github.com/f1nn-ach/pj-golang/initializers"
	"github.com/f1nn-ach/pj-golang/model"
)

func AddPets(pet model.Pet, email string) (sql.Result, error) {
	db := initializers.OpenConnection()
	defer db.Close()

	query := "INSERT INTO Pets VALUES (?,?,?,?,?,?,?)"
	result, err := db.Exec(query,
		pet.Id,
		pet.Name,
		pet.Gender,
		pet.Age,
		pet.Breed,
		pet.Species,
		email,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}
