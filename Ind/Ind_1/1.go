package main

import "fmt"

func main() {
	// Переменные для длины, ширины основания и высоты пирамиды
	var length, width, height float64

	// Вводим значения
	fmt.Println("Введите длину основания пирамиды:")
	fmt.Scan(&length)

	fmt.Println("Введите ширину основания пирамиды:")
	fmt.Scan(&width)

	fmt.Println("Введите высоту пирамиды:")
	fmt.Scan(&height)

	// Объем пирамиды
	volume := (length * width * height) / 3.0

	// езультат
	fmt.Printf("Объем пирамиды: %.2f\n", volume)
}
