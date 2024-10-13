package main

import (
	"fmt"
	"log"

	"github.com/f1nn-ach/pj-golang/controller"
	"github.com/f1nn-ach/pj-golang/initializers"
)

func init() {
	initializers.ConnectDatabase()
}

func main() {
	user, err := controller.GetUsers("jane.smith@example.com")

	if err != nil {
		log.Fatalf("Error retrieving user: %v", err)
	}

	fmt.Printf("User: %s\n", user.Email)
}
