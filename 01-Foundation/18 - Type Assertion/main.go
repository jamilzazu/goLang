package main

import "fmt"

func main() {
	var name interface{} = "Jamil Zazu"
	println("no type defined, name:", name)
	println("assigning string type, name:", name.(string))

	println("check if it is possible to assign", name)
	value, canAssigned := name.(int)
	fmt.Printf("the value is  %v, can be assigned = %v \n", value, canAssigned)

	name = 1
	println("no type defined, name:", name)
	println("assigning int type, name:", name.(int))
}
