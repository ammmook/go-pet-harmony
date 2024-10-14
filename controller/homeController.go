package controller

import (
	"fmt"
	"log"

	managementdb "github.com/f1nn-ach/pj-golang/managementDB"
)

func CallUser() {
	user, err := managementdb.GetUsersByEmail("user")

	if err != nil {
		log.Fatalf("Error retrieving user: %v", err)
	}

	fmt.Printf("User: %s\n", user.Email)
	fmt.Printf("Role: %s\n", user.Role)
}
