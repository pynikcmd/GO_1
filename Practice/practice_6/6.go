package main

import "fmt"

func main() {
	var number int

	// Ввод числа с клавиатуры
	fmt.Println("Введите натуральное число:")
	fmt.Scan(&number)

	// Находим последнюю цифру числа
	result := number % 10

	// Выводим последнюю цифру
	fmt.Println(result)
}
