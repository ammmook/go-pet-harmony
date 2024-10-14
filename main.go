package main

import (
	"fmt"
	"net/http"

	"github.com/f1nn-ach/pj-golang/controller"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "view/index.html")
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("view/net"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("view/assets"))))

	http.HandleFunc("/getUser", controller.CallUser)
	http.HandleFunc("/", GetIndex)

	http.HandleFunc("/login", controller.UserLogin)
	http.HandleFunc("/register", controller.UserRegister)

	http.HandleFunc("/result", controller.GetResultPage)

	fmt.Println("localhost:8000")
	http.ListenAndServe(":8000", nil)
}
