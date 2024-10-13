package main

type Booking struct {
	Id        string
	StartDate string
	EndDate   string
	Pet       Pet
	User      User
}

type Pet struct {
	Id      string
	Name    string
	Genter  string
	Age     string
	Breed   string
	Species string
	Booking []Booking
}

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

func main() {

}
