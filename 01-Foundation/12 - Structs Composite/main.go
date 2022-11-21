package main

import "fmt"

type Address struct {
	AddressName string
	Number      int
	City        string
	State       string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func main() {
	client := buildCClient()
	client.Address = buildCompositeAddress(client)
	RunPrint(client)
}

func buildCClient() Client {
	client := Client{
		Name:   "Jamil",
		Age:    34,
		Active: true,
	}
	return client
}

func buildCompositeAddress(client Client) Address {
	client.Address.AddressName = "Rua A" // or client.AddressName
	client.Address.Number = 10           // or client.Number
	client.Address.City = "São Paulo"    // or client.City
	client.Address.State = "SP"          // or client.State
	return client.Address
}

func RunPrint(client Client) {
	fmt.Printf("Name: %s \n", client.Name)
	fmt.Printf("Age: %d \n", client.Age)
	fmt.Printf("Active: %t \n", client.Active)
	fmt.Printf("Address: %s \n", client.Address.AddressName)
	fmt.Printf("Number: %d \n", client.Address.Number)
	fmt.Printf("City: %s \n", client.Address.City)
	fmt.Printf("State: %s \n", client.Address.State)
}
