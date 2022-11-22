package main

func main() {
	println("use simple for")
	for i := 0; i < 4; i++ {
		println(i)
	}

	println("\nuse simple for similar to while")
	i := 0
	for i < 2 {
		println(i)
		i++
	}

	numbers := []string{"one", "two", "three"}
	println("\nuse range")
	for key, value := range numbers {
		println(key, value)
	}

	println("\nonly values")
	for _, value := range numbers {
		println(value)
	}

	println("\ninfinite loop")
	for {
		println("infinite loop")
	}
}
