package main

import "fmt"

type Person interface {
	Deactivate()
}

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	client := BuildCClient()
	Deactivation(client)
}

func BuildCClient() Client {
	client := Client{
		Name:   "Jamil",
		Age:    34,
		Active: true,
	}
	return client
}

func Deactivation(person Person) {
	person.Deactivate()
}

func (client Client) Deactivate() {
	fmt.Println("Deactivating Client")
	client.Active = false
	fmt.Printf("Name: %s \n", client.Name)
	fmt.Printf("Age: %d \n", client.Age)
	fmt.Printf("Active: %t \n", client.Active)
}
