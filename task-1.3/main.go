package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := make([]int, 5)

	fmt.Println("Введите 5 натуральных чисел:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&numbers[i])
	}

	fmt.Printf("Вход: %d %d %d %d %d\n", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])

	sorted := make([]int, 5) // Сортируем копию по убыванию
	copy(sorted, numbers)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

	max := sorted[0] // самое большое - первый элемент отсортированного слайса
	min := sorted[4] // самое маленькое - последний элемент

	sum := 0
	for _, num := range numbers {
		sum += num
	}
	average := float64(sum) / 5

	fmt.Printf("Отсортированные элементы: %d %d %d %d %d\n", sorted[0], sorted[1], sorted[2], sorted[3], sorted[4])
	fmt.Printf("Самое большое число: %d\n", max)
	fmt.Printf("Самое маленькое число: %d\n", min)
	fmt.Printf("Среднее арифметическое: %.1f\n", average)
}
