package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

/*
Задача 3:
Создайте функцию capitalizeWords(s string) string,
которая преобразует каждое слово в строке так,
чтобы первая буква была заглавной, а остальные — строчными.
Например: "привет мир" → "Привет Мир".
*/

func main() {
	fmt.Println(`Программа 3
Введите строку:`)

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input := scanner.Text()
		output := capitalizeWords(input)
		fmt.Println(output)
	}
}

func capitalizeWords(s string) string {
	output := []rune{}
	// проходимся по каждому символу:
	// определяем начала слов по следованию за пробелом, ставим флаг flagForSpace,
	// далее всегда переводим flagForSpace в false.
	flagForSpace := true // начало строки, всегда с заглавной
	for _, char := range s {
		if unicode.IsSpace(char) {
			flagForSpace = true
			output = append(output, char)
		} else {
			if flagForSpace {
				output = append(output, unicode.ToUpper(char))
			} else {
				output = append(output, char)
			}
			flagForSpace = false
		}
	}
	return string(output)
}
