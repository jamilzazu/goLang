package main

func main() {
	firstNumber := 10
	secondNumber := 20

	useIf(firstNumber, secondNumber)
	useSwitch(firstNumber)

}

func useIf(firstNumber, secondNumber int) {
	if firstNumber > secondNumber {
		print(firstNumber)
	} else {
		println(secondNumber)
	}
}

func useSwitch(firstNumber int) {
	switch firstNumber {
	case 10:
		println(firstNumber)
	default:
		println("default")
	}
}
