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

	http.HandleFunc("/getUser", controller.CallUser)
	http.HandleFunc("/", controller.GetIndexPage)

	http.HandleFunc("/login", controller.UserLogin)
	http.HandleFunc("/register", controller.UserRegister)

	http.HandleFunc("/result", controller.GetResultPage)

	http.HandleFunc("/logout", controller.Logout)

	http.HandleFunc("/petregister", controller.PetRegister)
	http.HandleFunc("/listpets", controller.LoadListPetPage)

	fmt.Println("localhost:8000")
	http.ListenAndServe(":8000", nil)
}
