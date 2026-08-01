package main

import (
	"fmt"
	"math"
)

// task 1
type Person struct {
	name string
	age  int
}

// task 2
type Book struct {
	title  string
	author string
	pages  int
}

// task 3
type Car struct {
	brand string
	year  int
}

// task 4
type Rectangle struct {
	height float64
	width  float64
}

func rectangleArea(rectangle Rectangle) float64 { // функция для нахождение площадь
	return rectangle.height * rectangle.width
}

func rectanglePerimeter(rectangle Rectangle) float64 { // функция для нахождения периметра
	return 2 * (rectangle.height + rectangle.width)
}

// task 5
type Student struct {
	name  string
	grade int
}

// task 6
type Circle struct {
	radius float64
}

func circleArea(circle Circle) float64 { // функция для нахождение площади
	return circle.radius * circle.radius * (math.Pi)
}

func circleHeight(cicle Circle) float64 { // функция для нахождение длину окружности
	return 2 * math.Pi * cicle.radius
}

// task 7

type Sportcar struct {
	brand string
	speed int
}

func checkSpeed(sportcar Sportcar) string {
	if sportcar.speed > 100 {
		return "Слишком быстро"
	} else if sportcar.speed < 60 {
		return "Медленно"
	} else {
		return "Нормальная скорость"
	}
}

func main() {
	// task 1
	fmt.Println("<----- TASK 1 ----->")

	var person1 Person = Person{"Daniyar", 22}
	fmt.Printf("Name: %s, Age: %d\n", person1.name, person1.age)

	// task 2
	fmt.Println()
	fmt.Println("<----- TASK 2 ----->")

	var book1 Book = Book{
		"Парасат майданы",
		"Төлен Әбдік",
		89,
	}

	fmt.Printf("Title: %s, Author: %s, Pages: %d\n", book1.title, book1.author, book1.pages)

	// task 3
	fmt.Println()
	fmt.Println("<----- TASK 3 ----->")

	var car1 Car = Car{"Toyota", 2022}
	var p_car1 *Car = &car1

	fmt.Println("Brand:", car1.brand)
	fmt.Println("Year:", car1.year)

	p_car1.year = 2025
	fmt.Println("Year:", car1.year)

	// task 4
	fmt.Println()
	fmt.Println("<----- TASK 4 ----->")

	rectangle := Rectangle{6, 7}
	fmt.Println("Area:", rectangleArea(rectangle))
	fmt.Println("Perimeter:", rectanglePerimeter(rectangle))

	// task 5
	fmt.Println()
	fmt.Println("<----- TASK 5 ----->")

	student1 := Student{"Daniyar", 100}
	student2 := Student{"Danchik", 99}

	fmt.Printf("Студент %s получил %d баллов\n", student1.name, student1.grade)
	fmt.Printf("Студент %s получил %d баллов\n", student2.name, student2.grade)

	// task 6
	fmt.Println()
	fmt.Println("<----- TASK 6 ----->")

	circle1 := Circle{6.7}
	circle2 := Circle{9.7}

	fmt.Println("Area of circle1:", circleArea(circle1))
	fmt.Println("Height of circle1:", circleHeight(circle1))
	fmt.Println()
	fmt.Println("Area of circle2:", circleArea(circle2))
	fmt.Println("Height of circle2:", circleHeight(circle2))

	// task 7
	fmt.Println()
	fmt.Println("<----- TASK 7 ----->")

	sportcar1 := Sportcar{"Ferrari", 100}
	sportcar2 := Sportcar{"McLaren", 80}
	sportcar3 := Sportcar{"Porsche", 14}

	fmt.Printf("Speed of %s is %s\n", sportcar1.brand, checkSpeed(sportcar1))
	fmt.Printf("Speed of %s is %s\n", sportcar2.brand, checkSpeed(sportcar2))
	fmt.Printf("Speed of %s is %s\n", sportcar3.brand, checkSpeed(sportcar3))

}
