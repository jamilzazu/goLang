package main

import "fmt"

type NewInt int

type TypeNumber interface {
	~int | ~float64 // use ~ use to accept a custom type such as NewInt
}

func ConstrainedSum[T TypeNumber](values map[string]T) T {
	var sum T
	for _, value := range values {
		sum += value
	}
	return sum
}

func Sum[T int | float64](values map[string]T) T {
	var sum T
	for _, value := range values {
		sum += value
	}
	return sum
}

func getMapInt() map[string]int {
	return map[string]int{"Jamil": 1000, "Cristiano": 2000, "Flavio": 3000}
}

func getMapFloat() map[string]float64 {
	return map[string]float64{"Jamil": 1000.05, "Cristiano": 2000.10, "Flavio": 3000.20}
}

func getMapNewInt() map[string]NewInt {
	return map[string]NewInt{"Jamil": 1001, "Cristiano": 2001, "Flavio": 3001}
}

func main() {
	fmt.Println("\nUsing unconstrained sum")
	fmt.Printf("Int %d \n", Sum(getMapInt()))
	fmt.Printf("Float %.2f \n", Sum(getMapFloat()))

	fmt.Println("\nUsing constrained sum")
	fmt.Printf("Int %d \n", ConstrainedSum(getMapInt()))
	fmt.Printf("Float %.2f \n", ConstrainedSum(getMapFloat()))

	fmt.Println("\nUsing constrained sum with custom type")
	fmt.Printf("NewInt %d \n", ConstrainedSum(getMapNewInt()))

}
