package main

//run in terminal to generate the package
//go mod init github.com/golang/packages

import (
	"fmt"
	account "github.com/golang/packages/Account"
)

func main() {
	newSum := account.Sum(10, 20)
	fmt.Println("Result:", newSum)
}
