package main

import (
	"fmt"
	"net/http"

	"github.com/f1nn-ach/pj-golang/controller"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("view/net"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("view/assets"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("view/javascript"))))

	http.HandleFunc("/", controller.IndexPage)            // Index page
	http.HandleFunc("/login", controller.UserLogin)       // User login
	http.HandleFunc("/register", controller.UserRegister) // User registration (assuming UserRegister is implemented)
	http.HandleFunc("/result", controller.GetResultPage)  // Result page after login
	http.HandleFunc("/logout", controller.UserLogout)     // User logout
	http.HandleFunc("/getUser", controller.CallUser)

	http.HandleFunc("/petregister", controller.PetRegister)
	http.HandleFunc("/listpets", controller.LoadListPetPage)
	http.HandleFunc("/editmypet", controller.EditPet)
	http.HandleFunc("/deletepet", controller.DeletePet)

	fmt.Println("localhost:8000")
	http.ListenAndServe(":8000", nil)
}
