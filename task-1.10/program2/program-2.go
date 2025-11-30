package main

import (
	"fmt"
	"math"
)

/*
Задача 2:
Полиморфное отображение данных

Описание: Создайте интерфейс Shape с методом Area() float64.
Реализуйте структуры Circle и Rectangle, которые реализуют этот интерфейс.
Напишите функцию, которая принимает слайс фигур ([]Shape) и выводит площадь каждой фигуры.

Требования:

    Интерфейс Shape.
    Структуры Circle (с радиусом) и Rectangle (с длиной и шириной).
    Функция для вывода площадей всех фигур в слайсе.
*/

// Интерфейс Shape
type Shape interface {
	Area() float64
}

// Структура Circle
type Circle struct {
	Radius float64
}

// Метод Area для Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Структура Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Метод Area для Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Функция для вывода площадей всех фигур
func printAreas(shapes []Shape) {
	for i, shape := range shapes {
		area := shape.Area()
		fmt.Printf("%d. Площадь: %.2f\n", i+1, area)
	}
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5.0},
		Rectangle{Width: 4.0, Height: 6.0},
		Circle{Radius: 3.5},
		Rectangle{Width: 10.0, Height: 2.5},
		Circle{Radius: 7.0},
	}

	printAreas(shapes)
}
