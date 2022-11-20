package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {

	jamil := Client{
		Name:   "Jamil",
		Age:    34,
		Active: true,
	}

	fmt.Printf("Name: %s \n", jamil.Name)
	fmt.Printf("Age: %d \n", jamil.Age)
	fmt.Printf("Active: %t \n\n", jamil.Active)

	fmt.Println("Deactivate Jamil")
	jamil.Active = false

	fmt.Printf("Name: %s \n", jamil.Name)
	fmt.Printf("Age: %d \n", jamil.Age)
	fmt.Printf("Active: %t \n", jamil.Active)
}
