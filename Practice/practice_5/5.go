package main

import "fmt"

func main() {
	var number int

	fmt.Println("Введите целое число:")
	fmt.Scan(&number)

	fmt.Println(number * number)
}
