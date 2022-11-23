package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Client struct {
	Name string `json:"n"`
	Age  int    `json:"a"` //  to omit property use `json: "-"`
}

type Account struct {
	Number  int
	Balance float64
}

func useMarshal(account Account) {
	newJson, err := json.Marshal(account)
	messageError(err)
	fmt.Println("\nMarshal:")
	println(string(newJson))
}

func useUnmarshal(newJson []byte) {
	var account Account
	err := json.Unmarshal(newJson, &account)
	messageError(err)
	fmt.Println("\nUnmarshal:")
	fmt.Printf("Balance %.2f \n", account.Balance)
}

func useEncode(account Account) {
	fmt.Println("\nEncode:")
	err := json.NewEncoder(os.Stdout).Encode(account)
	messageError(err)
}

func useTags() {
	fmt.Println("\nTags:")
	jsonTag := []byte(`{"n":"Jamil","a":34}`)
	var client Client
	err := json.Unmarshal(jsonTag, &client)
	messageError(err)
	fmt.Println(client.Name)
}

func messageError(message error) {
	if message != nil {
		panic(message)
	}
}

func main() {
	account := Account{Number: 1, Balance: 200.03}
	json := []byte(`{"Number":1,"Balance":200.03}`)
	useMarshal(account)
	useUnmarshal(json)
	useEncode(account)
	useTags()
}
