package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	client := BuildCClient()
	RunPrint(client)
}

func BuildCClient() Client {
	client := Client{
		Name:   "Jamil",
		Age:    34,
		Active: true,
	}
	return client
}

func (client Client) Deactivate() Client {
	client.Active = false
	return client
}

func RunPrint(client Client) {
	fmt.Printf("Name: %s \n", client.Name)
	fmt.Printf("Age: %d \n", client.Age)
	fmt.Printf("Active: %t \n", client.Active)

	fmt.Println("Deactivating Client")
	client = client.Deactivate()

	fmt.Printf("Name: %s \n", client.Name)
	fmt.Printf("Age: %d \n", client.Age)
	fmt.Printf("Active: %t \n", client.Active)

}
