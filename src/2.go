package main

import (
	"fmt"
	"math"
)

func main() {
	var x int = 10
	var y int = 20
	z := math.Max(x, y)
	fmt.Println("Max value is", z)
}