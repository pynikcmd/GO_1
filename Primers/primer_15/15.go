package main

import "fmt"

const (
	c0 = iota // c0 == 0
	c1 = iota // c1 == 1
	c2 = iota // c2 == 2

)

func main() {
	fmt.Println(c0, c1, c2) // вывод: 0 1 2
}
