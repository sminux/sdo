package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Задача 3: Использование мапов для подсчета частот
Напишите программу, которая читает строку текста и подсчитывает количество вхождений каждого слова.
Используйте мапу (map[string]int) для хранения результатов.
Выведите полученную статистику.
*/

func main() {
	fmt.Println("Программа 3\nВведите строку:")

	scanner := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	if scanner.Scan() {
		input := scanner.Text()
		wordsArr := strings.Fields(input) // разбиваем на массив слов
		fmt.Println()
		for _, v := range wordsArr {
			words[v] = words[v] + 1 // храним счетчик слов в value мапы
		}
	}

	// отображаем слово (key) и подсчитанную статистику (value)
	fmt.Println("Слово\t: статистика использования (шт.)")
	for key, value := range words {
		fmt.Println(key, "\t:", value)
	}

}
