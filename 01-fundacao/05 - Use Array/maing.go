package main

import "fmt"

func main() {
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	for key, value := range myArray {
		fmt.Printf("O valor do indice é %d e o valor é %d \n", key, value)
	}
}
