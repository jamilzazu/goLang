package main

import "fmt"

func main() {

	fmt.Print("return func int value: ")
	newSumInt := sumInt(5, 10)
	fmt.Println(newSumInt)

	fmt.Print("return func float value: ")
	newSumFloat := sumFloat(10.2, 18.56)
	fmt.Println(newSumFloat)

	fmt.Print("return func string value: ")
	name := getName("Jamil")
	fmt.Println(name)

	fmt.Print("return func bool value: ")
	isTrue := isTrue(true)
	fmt.Println(isTrue)
}

func sumInt(startValue, endValue int) int {
	return startValue + endValue
}

func sumFloat(startValue, endValue float64) float64 {
	return startValue + endValue
}

func getName(name string) string {
	return name
}

func isTrue(isTrue bool) bool {
	return isTrue
}
