package math

type math struct {
	value1 int
	value2 int
}

func NewMath(value1, value2 int) math {
	return math{value1: value1, value2: value2}
}

func (math math) Add() int {
	return math.value1 + math.value2
}
