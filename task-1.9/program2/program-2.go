package main

import "fmt"

/*
Задача 2:
Структура "Студент" и метод для вычисления среднего балла
Описание:
Создайте структуру Student, которая содержит поля:
имя (Name) и список оценок (Grades []float64).
Реализуйте метод AverageGrade() float64, который возвращает средний балл студента.

Что нужно сделать:
Объявить структуру и метод.
Создать студента с несколькими оценками и вывести его средний балл.
*/

type Student struct {
	Name   string
	Grades []float64
}

func (s Student) AverageGrade() float64 {
	// Проверка на пустой список оценок
	if len(s.Grades) == 0 {
		return 0.0
	}

	// Проверка на минимальное количество оценок
	if len(s.Grades) < 3 {
		fmt.Printf("Внимание!\nу студента %s количество оценок: %d\n",
			s.Name, len(s.Grades))
	}

	sum := 0.0
	minGrade := s.Grades[0]
	maxGrade := s.Grades[0]

	for _, grade := range s.Grades {
		sum += grade
		if grade < minGrade {
			minGrade = grade
		}
		if grade > maxGrade {
			maxGrade = grade
		}
	}

	average := sum / float64(len(s.Grades))

	fmt.Printf("%s: мин=%g, макс=%g, средний=%.2f\n",
		s.Name, minGrade, maxGrade, average)

	return average
}

func main() {
	student1 := Student{
		Name:   "Дмитрий В.",
		Grades: []float64{4.5, 3.8, 5.0, 4.2},
	}

	student2 := Student{
		Name:   "Ульяна В.",
		Grades: []float64{5.0, 5.0}, // Мало оценок
	}

	student3 := Student{
		Name:   "Анатолий К.",
		Grades: []float64{}, // Нет оценок
	}

	fmt.Println("Расчет средних баллов:")
	student1.AverageGrade()
	student2.AverageGrade()
	student3.AverageGrade()
}
