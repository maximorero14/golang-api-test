package models

type Activity struct {
	Id     string //pk
	Type   string //sk
	UserId string
}
