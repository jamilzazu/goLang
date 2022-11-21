package main

import "fmt"

type Account struct {
	balance float64
}

func newAccount() *Account {
	return &Account{
		balance: 0.0,
	}
}

func (account *Account) income(value float64) float64 {
	fmt.Printf("\nBalance %.2f \n", account.balance)
	fmt.Printf("New Income %.2f \n", value)
	account.balance += value
	fmt.Printf("Total %.2f \n", account.balance)
	return account.balance
}

func main() {
	account := newAccount()
	account.income(350.0)
	account.income(550.50)
}
