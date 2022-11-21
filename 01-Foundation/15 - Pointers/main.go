package main

import "fmt"

func sumCopy(firstNumber, secondNumber int) int {
	fmt.Println("\nas it is a cloned value from memory, changing the variables does not interfere with the value stored in memory")
	firstNumber = 1000
	secondNumber = 1000
	return firstNumber + secondNumber
}

func sumMemory(firstNumber, secondNumber *int) int {
	fmt.Println("\nvalues changed through pointers, change the values stored in memory")
	*firstNumber = 1000
	*secondNumber = 500
	return *firstNumber + *secondNumber
}

func main() {
	fmt.Println("Memory -> Address -> Value")
	number := 10
	var pointer *int = &number
	fmt.Printf("number value is %d \n", number)
	fmt.Printf("pointer value is %p \n", pointer)

	fmt.Println("\nchanging number value")
	*pointer = 20
	fmt.Printf("number value is %d \n", number)
	fmt.Printf("pointer value is %p \n", pointer)

	fmt.Println("\nchanging number value")
	newNumber := &number
	*newNumber = 30
	fmt.Printf("number value is %d \n", number)
	fmt.Printf("pointer value is %p \n", pointer)

	firstNumber := 30
	secondNumber := 80

	sumCopy(firstNumber, secondNumber)
	fmt.Printf("firstNumber: %d, secondNumber: %d \n", firstNumber, secondNumber)

	sumMemory(&firstNumber, &secondNumber)
	fmt.Printf("firstNumber: %d, secondNumber: %d \n", firstNumber, secondNumber)
}
