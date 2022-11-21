package main

import "fmt"

type Person interface {
	deactivate()
}

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	client := buildCClient()
	deactivation(client)
}

func buildCClient() Client {
	client := Client{
		Name:   "Jamil",
		Age:    34,
		Active: true,
	}
	return client
}

func deactivation(person Person) {
	person.deactivate()
}

func (client Client) deactivate() {
	fmt.Println("Deactivating Client")
	client.Active = false
	fmt.Printf("Name: %s \n", client.Name)
	fmt.Printf("Age: %d \n", client.Age)
	fmt.Printf("Active: %t \n", client.Active)
}
