package main

import (
	"fmt"
	"math"
)

func main() {
	isNegative := math.Signbit(-5)
	fmt.Println(isNegative)
}
