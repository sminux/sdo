package main

import (
	"fmt"
	"os"
	"slices"

	"golang.org/x/term"
)

/*
Задача 2: Манипуляции со слайсами
Создайте слайс строк, содержащий названия городов.
Реализуйте функции для добавления нового города,
удаления города по имени и поиска города в списке.
Продемонстрируйте работу этих функций на примере.
*/

func main() {
	fmt.Println("Программа 2")

	cities := []string{"Ак-Довурак", "Аксай", "Алагир", "Алапаевск", "Алатырь",
		"Алдан", "Алейск", "Александров", "Александровск", "Александровск-Сахалинский",
		"Алексеевка", "Алексин", "Алзамай", "Алупка", "Алушта",
		"Альметьевск", "Амурск", "Анадырь", "Анапа", "Ангарск",
		"Андреаполь", "Анжеро-Судженск", "Анива", "Апатиты", "Апрелевка",
		"Апшеронск", "Арамиль", "Аргун", "Ардатов", "Ардон",
		"Арзамас", "Аркадак", "Армавир", "Армянск", "Арсеньев",
		"Арск", "Артём", "Артёмовск", "Артёмовский", "Архангельск",
		"Асбест", "Асино", "Астрахань", "Аткарск", "Ахтубинск",
		"Ачинск", "Аша"}

	// Добавление города
	addCity(&cities, "Москва")
	fmt.Println("Список городов после добавления Москвы: \n", cities)

	printFullWidthSeparator("-")

	// Удаление города
	delCity(&cities, "Ак-Довурак")
	fmt.Println("Список городов после удаления Ак-Довурака: \n", cities)

	printFullWidthSeparator("-")

	// Поиск города
	searchCity(&cities, "Армянск")
}

func addCity(cities *[]string, newCity string) {
	*cities = append(*cities, newCity)
}

func delCity(cities *[]string, delCity string) {
	// https://cs.opensource.google/go/go/+/go1.25.4:src/slices/slices.go;l=96
	indexToRemove := slices.Index(*cities, delCity)
	if indexToRemove == -1 {
		return
	}
	// SliceTricks https://go.dev/wiki/SliceTricks#delete
	*cities = append((*cities)[:indexToRemove], (*cities)[indexToRemove+1:]...)
}

func searchCity(cities *[]string, delCity string) {
	index := slices.Index(*cities, delCity)
	if index == -1 {
		return
	}
	fmt.Printf("Индекс города %s - %v\n", delCity, index)
}

// дополнительные функции для разделения результатов
func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80 // стандартная ширина по умолчанию
	}
	return width
}

func printFullWidthSeparator(char string) {
	width := getTerminalWidth()
	for i := 0; i < width; i++ {
		fmt.Print(char)
	}
	fmt.Println()
}
