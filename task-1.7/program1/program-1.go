package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
Задача 1: Работа с массивами и слайсами
Напишите программу, которая создает массив из 10 целых чисел, заполняет его случайными значениями от 1 до 100.
Затем скопируйте этот массив в слайс и отсортируйте его по возрастанию.
Выведите исходный массив и отсортированный слайс.
*/

func main() {
	fmt.Println("Программа 1")

	// https://www.geeksforgeeks.org/go-language/generating-random-numbers-in-golang/
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)

	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = y1.Intn(100)
	}

	// Создаем слайс и копируем в него массив
	slice := make([]int, len(arr))
	copy(slice, arr[:])

	// Сортируем слайс используя sort
	sort.Ints(slice)

	fmt.Printf("Исходный массив: %v\n", arr)
	fmt.Printf("Отсортированный слайс: %v\n", slice)
}
