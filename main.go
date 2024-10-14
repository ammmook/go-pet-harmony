package main

import (
	"github.com/f1nn-ach/pj-golang/controller"
	"github.com/f1nn-ach/pj-golang/initializers"
)

func init() {
	initializers.ConnectDatabase()
}

func main() {
	controller.CallUser()
}
