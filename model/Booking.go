package model

type Booking struct {
	Id        string
	StartDate string
	EndDate   string
	Pet       Pet
	User      User
}
