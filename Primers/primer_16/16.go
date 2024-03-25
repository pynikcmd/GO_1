package main

import "fmt"

const (
	u         = iota * 42 // u == 0 (индекс - 0, поэтому 0 * 42 = 0)
	v float64 = iota * 42 // v == 42.0 (индекс - 1, поэтому 1.0 * 42 = 42.0)
	w         = iota * 42 // w == 84 (индекс - 2, поэтому 2 * 42 = 84)

)

// переменные ни в одноме блоке const, поэтому индекс не увеличился
const x = iota // x == 0
const y = iota // y == 0

func main() {
	fmt.Println(x, y)    // вывод: 0 0
	fmt.Println(u, v, w) // вывод: 0 42 84
}
