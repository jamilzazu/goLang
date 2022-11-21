package main

import "fmt"

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
}
