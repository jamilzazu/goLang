package account

type TypeNumber interface {
	~int | ~float64 // use ~ use to accept a custom type such as NewInt
}

func Sum[T TypeNumber](firstValues, secondValue T) T {
	return firstValues + secondValue
}
