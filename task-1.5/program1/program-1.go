package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

/*
Задача 1:
Напишите программу, которая запрашивает у пользователя ввод строки,
а затем выводит число - количество символов в строке
*/

func main() {
	fmt.Println("Программа 1")
	fmt.Println("Введите строку:")

	// Scanln() неравильно считывает - только до пробела
	// var input string
	// fmt.Scanln(&input)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		// len() неправильно работает с кириллицей (считает один символ за 2 байта)
		// fmt.Println("Количество символов в строке: ", len(input))

		runeLength := utf8.RuneCountInString(input)
		fmt.Println("Количество символов в строке: ", runeLength)
	}

}
