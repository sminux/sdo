package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
Задача 1:
Напишите программу, которая принимает список имен файлов в текущей директории,
объединяет их содержимое и сохраняет результат в новый файл combined.txt (работать только с текстовыми файлами)
*/

func getTextFiles(path string) ([]string, error) {
	entries, err := os.ReadDir(path) //  текущая директория
	if err != nil {
		return nil, err
	}

	var textFiles []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()
		ext := strings.ToLower(filepath.Ext(filename))

		// Список текстовых расширений
		textExtensions := map[string]bool{
			".txt":  true,
			".log":  true,
			".md":   true,
			".json": true,
			".xml":  true,
			".yaml": true,
			".yml":  true,
			".csv":  true,
		}

		if textExtensions[ext] {
			textFiles = append(textFiles, filename)
		}
	}

	return textFiles, nil
}

func combineTextFiles(filenames []string, output string) error {
	out, err := os.Create(output)
	if err != nil {
		return err
	}
	defer out.Close()

	writer := bufio.NewWriter(out)
	defer writer.Flush()

	totalFiles := 0

	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil { // пропускаем файл, если нет прав на чтение
			fmt.Printf("Ошибка чтения %s: %v\n", filename, err)
			continue
		}
		defer file.Close()

		// Копируем содержимое
		scanner := bufio.NewScanner(file)
		lineCount := 0

		for scanner.Scan() {
			if _, err := writer.WriteString(scanner.Text() + "\n"); err != nil {
				return err
			}
			lineCount++
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Ошибка чтения %s: %v\n", filename, err)
			continue
		}

		totalFiles++
	}

	if totalFiles == 0 {
		return fmt.Errorf("не обработано ни одного файла")
	}

	return nil
}

func main() {
	fmt.Println("Задача 2")
	const path = "."

	textFiles, err := getTextFiles(path)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	if len(textFiles) == 0 {
		fmt.Println("Текстовые файлы не найдены")
		return
	}

	fmt.Println("Найденные текстовые файлы:")
	for i, file := range textFiles {
		fmt.Printf("%d) %s\n", i+1, file)
	}

	outputFile := "combined.txt"

	if err := combineTextFiles(textFiles, outputFile); err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Printf("Результат: %s\n", outputFile)
}
