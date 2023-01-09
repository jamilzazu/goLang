package main

import (
	"fmt"
	"github.com/jamilzazu/golang/05/01/math"
)

func main() {
	// go mod ini
	sum := math.NewMath(1, 2)
	fmt.Println(sum.Add())
}
