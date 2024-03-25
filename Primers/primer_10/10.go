package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("Введите три целых числа через пробел:")

	fmt.Scan(&a, &b, &c)

	fmt.Println("Вы ввели:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
}
