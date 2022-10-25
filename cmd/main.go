package main

import "fmt"

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
}
