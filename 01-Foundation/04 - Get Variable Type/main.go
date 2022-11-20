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
	fmt.Printf("the type of booleanVariable is %T \n", booleanVariable)
	fmt.Printf("the type of intVariable is %T \n", intVariable)
	fmt.Printf("the type of stringVariable is %T \n", stringVariable)
	fmt.Printf("the type of floatVariable is %T \n", floatVariable)
	fmt.Printf("the type of newType is %T \n", nType)
}
