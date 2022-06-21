package models

type Activity struct {
	ID     string //pk
	Type   string //sk
	UserId string
}
