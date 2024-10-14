package main

import (
	"net/http"

	"github.com/f1nn-ach/pj-golang/controller"
	"github.com/f1nn-ach/pj-golang/initializers"
)

func init() {
	initializers.ConnectDatabase()
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("net"))))
	controller.CallUser()
}
