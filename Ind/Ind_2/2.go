package main

import (
	"fmt"
	"math"
)

func main() {
	// Задаем значения катетов прямоугольного треугольника
	var a, b float64
	fmt.Println("Введите длину первого катета:")
	fmt.Scan(&a)
	fmt.Println("Введите длину второго катета:")
	fmt.Scan(&b)

	// Вычисляем гипотенузу
	c := math.Sqrt(a*a + b*b)

	// Вычисляем периметр
	per := a + b + c

	// Выводим результат
	fmt.Printf("Периметр прямоугольного треугольника: %.2f\n", per)
}
