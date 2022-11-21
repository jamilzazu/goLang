package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	client := buildCClient()
	runPrint(client)
}

func buildCClient() Client {
	client := Client{
		Name:   "Jamil",
		Age:    34,
		Active: true,
	}
	return client
}

func (client Client) deactivate() Client {
	client.Active = false
	return client
}

func runPrint(client Client) {
	fmt.Printf("Name: %s \n", client.Name)
	fmt.Printf("Age: %d \n", client.Age)
	fmt.Printf("Active: %t \n", client.Active)

	fmt.Println("Deactivating Client")
	client = client.deactivate()

	fmt.Printf("Name: %s \n", client.Name)
	fmt.Printf("Age: %d \n", client.Age)
	fmt.Printf("Active: %t \n", client.Active)

}
