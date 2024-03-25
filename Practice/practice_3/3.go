package main

import "fmt"

func main() {
	var number int

	fmt.Println("Введите число:")
	fmt.Scan(&number)

	result := number * 2
	result += 100

	fmt.Println("Результат:", result)
}
