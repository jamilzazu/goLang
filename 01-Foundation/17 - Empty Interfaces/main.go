package main

import "fmt"

func main() {
	var sector interface{} = 10
	var tank interface{} = "387A"

	fmt.Println("Empty interfaces accept any value, like Generics")

	changeNameSector(sector)
	changeNameTank(tank)
}

func changeNameSector(value interface{}) {
	fmt.Printf("\nOld sector name: %v, old type: %T \n", value, value)
	value = "Sector 10"
	fmt.Printf("New sector name: %v, new type: %T \n", value, value)
}

func changeNameTank(value interface{}) {
	fmt.Printf(" \nOld tank name: %v, old type: %T \n", value, value)
	value = 581
	fmt.Printf("New tank name: %v, new type: %T \n", value, value)
}
