package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

/*
Задача 2:
Напишите программу, которая подсчитывает количество
гласных букв (а, е, ё, и, о, у, ы, э, ю, я) в введённой пользователем строке.
*/

func main() {
	fmt.Println("Программа 2")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите строку: ")
	if scanner.Scan() {
		input := scanner.Text()
		vowelCount, _, _ := analyzeRussianLetters(input)

		fmt.Printf("\nАнализ строки: \"%s\"\n", input)
		fmt.Printf("Гласные буквы: %d\n", vowelCount)
		// для отладки и проверки:
		// fmt.Printf("Согласные буквы: %d\n", consonantCount)
		// fmt.Printf("Другие символы: %d\n", otherCount)
		// fmt.Printf("Всего символов: %d\n", len([]rune(input)))
	}
}

func analyzeRussianLetters(s string) (vowels, consonants, others int) {
	russianVowels := map[rune]bool{
		'а': true, 'е': true, 'ё': true, 'и': true, 'о': true,
		'у': true, 'ы': true, 'э': true, 'ю': true, 'я': true,
	}

	russianConsonants := map[rune]bool{
		'б': true, 'в': true, 'г': true, 'д': true, 'ж': true, 'з': true,
		'й': true, 'к': true, 'л': true, 'м': true, 'н': true, 'п': true,
		'р': true, 'с': true, 'т': true, 'ф': true, 'х': true, 'ц': true,
		'ч': true, 'ш': true, 'щ': true,
	}

	for _, char := range s {
		lowerChar := unicode.ToLower(char)

		if russianVowels[lowerChar] {
			vowels++
		} else if russianConsonants[lowerChar] {
			consonants++
		} else {
			others++
		}
	}

	return vowels, consonants, others
}
