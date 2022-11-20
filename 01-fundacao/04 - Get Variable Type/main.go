package main

import (
	"fmt"
)

type newType int

var (
	booleanVariable         = true
	intVariable             = 10
	stringVariable          = "Jamil"
	floatVariable           = 1.2
	nType           newType = 1000
)

func main() {
	fmt.Printf("O tipo de booleanVariable é %T \n", booleanVariable)
	fmt.Printf("O tipo de intVariable é %T \n", intVariable)
	fmt.Printf("O tipo de stringVariable é %T \n", stringVariable)
	fmt.Printf("O tipo de floatVariable é %T \n", floatVariable)
	fmt.Printf("O tipo de newType é %T \n", nType)
}
