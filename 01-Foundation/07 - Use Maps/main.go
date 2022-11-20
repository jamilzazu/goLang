package main

import "fmt"

func main() {
	salaries := map[string]int{"Jamil": 1000, "Cristiano": 2000, "Flavio": 3000, "Miguel": 4000}

	for name, salary := range salaries {
		fmt.Printf("%s's salary is %d \n", name, salary)
	}

	fmt.Printf("\nremoving jamil's salary \n")
	delete(salaries, "Jamil")

	for name, salary := range salaries {
		fmt.Printf("%s's salary is %d \n", name, salary)
	}

	fmt.Printf("\nraising Cristiano's salary \n")
	salaries["Cristiano"] = 25000

	for name, salary := range salaries {
		fmt.Printf("%s's salary is %d \n", name, salary)
	}

	fmt.Printf("\nwages only \n")
	for _, salary := range salaries {
		fmt.Printf("Salary is %d\n", salary)
	}
}
