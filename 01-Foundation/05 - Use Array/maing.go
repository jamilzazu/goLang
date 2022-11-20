package main

import "fmt"

func main() {
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	for key, value := range myArray {
		fmt.Printf("The index is %d the value is %d \n", key, value)
	}
}
