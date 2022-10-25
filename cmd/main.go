package main

import (
	"fmt"

	"github.com/kapralovs/git_basics/internal/models"
)

func main() {
	fmt.Println("hello world!")
	someUser := models.User{
		ID:      1,
		Name:    "John",
		Surname: "Doe",
		Age:     42,
		Weight:  81.7,
	}

	fmt.Println(someUser)

	vw := new(models.Car)
	vw.Model = "Volkswagen Polo"
	someUser.Vehicle = vw

	fmt.Println(someUser.Vehicle.Model)
}
