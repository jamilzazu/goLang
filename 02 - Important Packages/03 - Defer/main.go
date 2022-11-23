package main

import (
	"fmt"
)

func returnSecondThirdFirstLine() {
	defer fmt.Println("first line")
	fmt.Println("second line")
	fmt.Println("third line")
}

func main() {
	fmt.Println("\ndefer execute command at end of method")
	returnSecondThirdFirstLine()
}
