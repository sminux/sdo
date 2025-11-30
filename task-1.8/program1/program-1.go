package main

import (
	"errors"
	"fmt"
)

/*
Задача 1:
Реализуйте функцию divide(a, b float64) (float64, error), которая делит a на b.
Если b равно нулю, возвращайте ошибку.
В основной программе вызовите эту функцию и обработайте возможную ошибку.
*/

func main() {
	fmt.Println("Программа 1")

	fmt.Println("Введите число a:")
	var a float64
	fmt.Scanln(&a)

	fmt.Println("Введите число b:")
	var b float64
	fmt.Scanln(&b)

	res, err := divide(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Результат:", res)

}

func divide(a, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("Ошибка: b равно нулю")
	}
	res := a / b
	return res, nil
}
