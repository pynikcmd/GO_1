package main

import "fmt"

func main() {
	var degrees int

	// Ввод числа градусов с клавиатуры
	fmt.Println("Введите число градусов:")
	fmt.Scan(&degrees)

	// Переводим градусы в часы и минуты
	hours := degrees / 30
	minutes := (degrees % 30) * 2

	// Выводим результат
	fmt.Printf("It is %d hours %d minutes.\n", hours, minutes)
}
