package managementdb

import (
	"database/sql"
	"fmt"

	"github.com/f1nn-ach/pj-golang/initializers"
	"github.com/f1nn-ach/pj-golang/model"
)

func AddPets(pet model.Pet, email string) (sql.Result, error) {
	db := initializers.OpenConnection()
	defer db.Close()

	query := "INSERT INTO Pets (name,gender,age,breed,species,user_email) VALUES (?,?,?,?,?,?)"
	result, err := db.Exec(query,
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

func GetPetsByEmail(email string) ([]model.Pet, error) {
	db := initializers.OpenConnection()
	defer db.Close()

	query := "SELECT * FROM Pets WHERE user_email=?"

	rows, err := db.Query(query, email)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var pets []model.Pet

	for rows.Next() {
		var pet model.Pet
		err := rows.Scan(&pet.Id, &pet.Name, &pet.Gender, &pet.Age, &pet.Breed, &pet.Species, &email)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		pets = append(pets, pet)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered while iterating rows: %w", err)
	}

	return pets, nil
}

func GetPetById(id int) (*model.Pet, error) {
	db := initializers.OpenConnection()
	defer db.Close()

	query := `
		SELECT p.id, p.name, p.gender, p.age, p.breed, p.species, u.email
		FROM Pets p
		LEFT JOIN Users u ON p.user_email = u.email
		WHERE p.id=?
	`

	result := db.QueryRow(query, id)

	var pet model.Pet
	var user model.User
	err := result.Scan(&pet.Id, &pet.Name, &pet.Gender, &pet.Age, &pet.Breed, &pet.Species, &user.Email)
	if err != nil {
		return nil, err
	}
	pet.Email = user
	return &pet, nil
}

func EditPets(pet model.Pet) (sql.Result, error) {
	db := initializers.OpenConnection()
	defer db.Close()

	query := `UPDATE Pets SET name=?, gender=?, age=?, breed=?, species=? WHERE id=?`
	result, err := db.Exec(query,
		pet.Name,
		pet.Gender,
		pet.Age,
		pet.Breed,
		pet.Species,
		pet.Id,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update pet with ID %d: %v", pet.Id, err)
	}
	return result, nil
}
