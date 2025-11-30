package main

import "fmt"

/*
Задача 2:
Функция высшего порядка: передача функции как аргумента
Создайте функцию applyOperation(a, b int, op func(int, int) int) int,
которая применяет переданную функцию op к числам a и b.
*/

func main() {
	fmt.Println("Программа 1")

	fmt.Println("Введите число a:")
	var a float64
	fmt.Scanln(&a)

	fmt.Println("Введите число b:")
	var b float64
	fmt.Scanln(&b)

	result1 := applyOperation(a, b, func(x, y float64) float64 {
		return x + y
	})
	fmt.Printf("%v + %v = %v\n", a, b, result1)

	result2 := applyOperation(a, b, func(x, y float64) float64 {
		return x - y
	})
	fmt.Printf("%v - %v = %v\n", a, b, result2)

	result3 := applyOperation(a, b, func(x, y float64) float64 {
		return x * y
	})
	fmt.Printf("%v - %v = %v\n", a, b, result3)
}

func applyOperation(a, b float64, op func(float64, float64) float64) float64 {
	return op(a, b)
}
