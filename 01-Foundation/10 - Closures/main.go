package main

import "fmt"

func main() {
	fmt.Println("Closures, calling a function inside another function")
	name := func() string {
		return getName()
	}()
	fmt.Println(name)
}

func getName() string {
	return "Jamil"
}
