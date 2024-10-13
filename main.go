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
	user, _ := controller.GetUsers(5)

	if user != nil {
		fmt.Printf("User : %v\n", user.Email)
	} else {
		log.Fatal("Can't get user")
	}
}
