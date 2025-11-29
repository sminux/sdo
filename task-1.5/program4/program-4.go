package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

/*
Задача 4:
Напишите программу, которая запрашивает у пользователя ввод строки-формулы,
а выводит сообщение о правильности написания круглых скобок, например:

Пример 1
строка на вход: (1+1)*(2+2)
вывод: Скобки расставлены верно, 2 открывающиеся, 2 закрывающиеся

Пример 2
Строка на вход: ((1+1) + (2+2) ))
вывод: Скобки расставлены неправильно, 3 открывающиеся, 4 закрывающиеся
*/

func main() {
	fmt.Println("Программа 4\nВведите строку:")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input := scanner.Text()
		output, err := checkBrackets(input)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(output)
	}
}

func checkBrackets(s string) (string, error) {
	var output string
	var countOpen, countClose int
	openFlag := false // проверяем, что первая скобка будет открывающей
	for _, char := range s {
		switch char {
		case '(':
			if !openFlag {
				openFlag = true
			}
			countOpen += 1
		case ')':
			if !openFlag {
				err := errors.New("Ошибка: первая скобка закрывающая")
				return "", err
			}
			countClose += 1
		default:
			// pass
		}

	}

	if countOpen == countClose {
		output = fmt.Sprintf("вывод: Скобки расставлены верно, %d открывающиеся, %d закрывающиеся\n", countOpen, countClose)
	} else {
		output = fmt.Sprintf("вывод: Скобки расставлены неправильно, %d открывающиеся, %d закрывающиеся\n", countOpen, countClose)
	}

	return output, nil
}
