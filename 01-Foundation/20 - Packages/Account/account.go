package account

import "fmt"

type Account struct {
	Limit float64
}

type TypeNumber interface {
	~int | ~float64 // use ~ use to accept a custom type such as NewInt
}

func Sum[T TypeNumber](firstValues, secondValue T) T {
	return firstValues + secondValue
}

func MessageMyBalance() {
	fmt.Println("Your balance is: ", 1250.99)
}

var InitialBalance = 100.00
