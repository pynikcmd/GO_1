package main

import "fmt"

func main() {
	var x int
	var y float64
	var max, min int
	x = 10
	var c string = "Hello World!"
	var z float64 = 1.045
	var a = 12
	var hello = "Hello"

	fmt.Printf("x: %v, type: %T\n", x, x)
	fmt.Printf("y: %v, type: %T\n", y, y)
	fmt.Printf("max: %v, type: %T\n", max, max)
	fmt.Printf("min: %v, type: %T\n", min, min)
	fmt.Printf("c: %v, type: %T\n", c, c)
	fmt.Printf("z: %v, type: %T\n", z, z)
	fmt.Printf("a: %v, type: %T\n", a, a)
	fmt.Printf("hello: %v, type: %T\n", hello, hello)
}
