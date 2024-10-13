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
	user, err := controller.GetUser(5)

	if err != nil {
		log.Fatal("Can't get user")
	} else {
		fmt.Printf("User email : %v\n", user.Email)
	}
}
