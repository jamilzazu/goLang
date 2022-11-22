package main

func main() {
	useSimpleFor()
	useRangeFor()
	useInfiniteFor()
}

func useSimpleFor() {
	println("use simple for")
	for i := 0; i < 4; i++ {
		println(i)
	}
}

func useRangeFor() {
	numbers := []string{"one", "two", "three"}
	println("\nuse range")
	for key, value := range numbers {
		println(key, value)
	}

	println("\nonly values")
	for _, value := range numbers {
		println(value)
	}
}

func useInfiniteFor() {
	println("\ninfinite loop")
	for {
		println("infinite loop")
	}
}
