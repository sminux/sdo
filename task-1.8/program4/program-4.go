package main

import "fmt"

/*
Задание 4:
Создайте функцию sumAll, которая принимает произвольное количество целых чисел и возвращает их сумму.
Пример использования:
    fmt.Println(sumAll(1, 2, 3)) // 6
    fmt.Println(sumAll(10, -2, 4, 7)) // 19
*/

func sumAll(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(sumAll(1, 2, 3))      // 6
	fmt.Println(sumAll(10, -2, 4, 7)) // 19
	fmt.Println(sumAll())             // 0
	fmt.Println(sumAll(5))            // 5
}
