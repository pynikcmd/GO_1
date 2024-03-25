package main

import "fmt"

func main() {
	a := 100
	b := 10
	c := a + b
	fmt.Println("c =", c) // Вывод: c = 110

	c = a * b
	fmt.Println("c =", c) // Вывод: c = 1000

	c = a - b
	fmt.Println("c =", c) // Вывод: c = 90

	c = a / b
	fmt.Println("c =", c) // Вывод: c = 10

	c = a % 3
	fmt.Println("c =", c) // Вывод: c = 1

	var d int = 1
	d++
	fmt.Println("d =", d) // Вывод: d = 2

	var e int = 10
	e--
	fmt.Println("e =", e) // Вывод: e = 9
}
