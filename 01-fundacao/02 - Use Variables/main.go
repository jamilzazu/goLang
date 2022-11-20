package main

var (
	defaultBooleanVariable bool
	defaultIntVariable     int
	defaultStringVariable  string
	defaultFloatVariable   float64

	booleanVariable = true
	intVariable     = 10
	stringVariable  = "Jamil"
	floatVariable   = 1.2
)

func main() {
	println("Variables Default Values")
	println(defaultBooleanVariable)
	println(defaultIntVariable)
	println(defaultStringVariable)
	println(defaultFloatVariable)
	println("")
	println("")
	println("Initialized Variables")
	println(booleanVariable)
	println(intVariable)
	println(stringVariable)
	println(floatVariable)

}
