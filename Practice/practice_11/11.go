package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Задаем значения полуосей эллипса
	var a, b float64
	fmt.Println("Введите значение a:")
	fmt.Scan(&a)
	fmt.Println("Введите значение b:")
	fmt.Scan(&b)

	// Количество случайных точек
	numPoints := 10000

	// Переменные для подсчета точек внутри эллипса
	pointsInside := 0

	// Генерируем случайные точки и подсчитываем количество точек внутри эллипса
	for i := 0; i < numPoints; i++ {
		x := rand.Float64()*a*2 - a
		y := rand.Float64()*b*2 - b
		if x*x/(a*a)+y*y/(b*b) <= 1 {
			pointsInside++
		}
	}

	// Оцениваем площадь поверхности эллипса
	surfaceArea := float64(pointsInside) / float64(numPoints) * (4 * a * b)

	// Оцениваем объем тела, образованного вращением эллипса вокруг оси Ox
	volume := surfaceArea * (2.0 / 3.0) * a

	// Выводим результат
	fmt.Printf("Площадь поверхности: %.2f\n", surfaceArea)
	fmt.Printf("Оценка объема тела: %.2f\n", volume)
}
