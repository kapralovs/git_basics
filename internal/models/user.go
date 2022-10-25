package models

type User struct {
	ID      int
	Name    string
	Surname string
	Age     int
	Weight  float32
	Vehicle *Car
}
