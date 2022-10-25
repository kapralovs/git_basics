package main

import (
	"fmt"

	"github.com/kapralovs/git_basics/internal/models"
)

func main() {
	fmt.Println("hello world with modules!")
	someUser := models.User{}

	fmt.Println(someUser)

	vw := new(models.Car)
	vw.Model = "Volkswagen Polo"
	someUser.Vehicle = vw

	fmt.Println(someUser.Vehicle.Model)
}
