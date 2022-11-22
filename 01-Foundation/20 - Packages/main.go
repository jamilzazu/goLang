package main

// run in terminal to generate the package
// go mod init github.com/golang/packages
// to make method, function, struct and variable public among the packages, the name must be capitalized. Otherwise, it will be private
// example function Account
import (
	"fmt"
	account "github.com/golang/packages/Account"
	"github.com/google/uuid" // run mod tidy to download external packages
)

func main() {
	newSum := account.Sum(10, 20)
	fmt.Println("Result:", newSum)
	fmt.Println("InitialBalance:", account.InitialBalance)

	limit := account.Account{Limit: 248.01}
	fmt.Println("Account.Limit:", limit)

	account.MessageMyBalance()

	fmt.Println(uuid.New())
}
