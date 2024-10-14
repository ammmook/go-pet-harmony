package controller

import (
	"fmt"
	"log"

	"github.com/f1nn-ach/pj-golang/initializers"
)

func CallUser() {
	user, err := initializers.GetUsers("user")

	if err != nil {
		log.Fatalf("Error retrieving user: %v", err)
	}

	fmt.Printf("User: %s\n", user.Email)
	fmt.Printf("Role: %s\n", user.Role)
}
