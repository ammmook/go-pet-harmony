package model

type User struct {
	Email       string
	Firstname   string
	Lastname    string
	PhoneNumber string
	Password    string
	Role        string
	Pet         []Pet
	Booking     []Booking
}
