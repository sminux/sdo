package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Задача 1:
Создайте программу, которая читает лог-файл server.log, подсчитывает количество строк, содержащих слово "error", и выводит это число.
*/

func countErrors(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("ошибка открытия файла: %v", err)
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), "error") {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("ошибка чтения файла: %v", err)
	}

	return count, nil
}

func main() {
	fmt.Println("Задача 1")
	filename := "server.log"

	count, err := countErrors(filename)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Printf("Количество строк содержащих 'error': %d\n", count)
}
