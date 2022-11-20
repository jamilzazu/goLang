package main

import "fmt"

func main() {
	newSlice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v \n", len(newSlice), cap(newSlice), newSlice)

	fmt.Printf("\n#get the first 0 values \n")
	fmt.Printf("len=%d cap=%d %v \n", len(newSlice[:0]), cap(newSlice[:0]), newSlice[:0])

	fmt.Printf("\n#get the first 4 values \n")
	fmt.Printf("len=%d cap=%d %v \n", len(newSlice[:4]), cap(newSlice[:4]), newSlice[:4])

	fmt.Printf("\n#ignore the first 2 values \n")
	fmt.Printf("len=%d cap=%d %v \n", len(newSlice[2:]), cap(newSlice[2:]), newSlice[2:])

	fmt.Printf("\n#adding new value: 110 \n")
	newSlice = append(newSlice, 110)
	fmt.Printf("\n#when adding a new value, the slice doubles the initial capacity \n")
	fmt.Printf("len=%d cap=%d %v\n", len(newSlice), cap(newSlice), newSlice)
}
