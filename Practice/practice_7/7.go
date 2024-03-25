package main

import "fmt"

func main() {
	var number int

	// Ввод числа с клавиатуры
	fmt.Println("Введите неотрицательное целое число:")
	fmt.Scan(&number)

	// Находим число десятков
	result := (number / 10) % 10

	// Выводим число десятков
	fmt.Println(result)
}
