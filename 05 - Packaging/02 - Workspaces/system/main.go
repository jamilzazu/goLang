package main

import (
	"github.com/google/uuid"
	"github.com/jamilzazu/golang/05-Packaging/02-Workspaces/math"
)

func main() {
	// Ignore dependencies in error
	// go mod tidy -e
	m := math.NewMath(1, 2)
	println(m.Add())
	println(uuid.New().String())
}
