package main

import "fmt"

func main() {
	fmt.Println(sum(1, 3, 45, 6, 34, 654, 654, 7645, 534, 543, 543, 543))
}

func sum(numbers ...int) int {
	fmt.Println("1, 3, 45, 6, 34, 654, 654, 7645, 534, 543, 543, 543")
	fmt.Println("read the parameters sent, as long as they are of the same type")

	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
