package main

import "fmt"

func main() {
	// task 1
	fmt.Println("------ Task 1 ---")

	schooling := 11
	fmt.Println("Schooling:", schooling)

	schooling = schooling + 1
	fmt.Println("Schooling:", schooling)

	// task 2
	fmt.Println()
	fmt.Println("------ Task 2 ---")

	var name string = "Vladislav"
	fmt.Println("Name:", name)

	name = "Daniyar"
	fmt.Println("New Name:", name)

	// task 3
	fmt.Println()
	fmt.Println("------ Task 3 ---")

	var steps int = 0
	fmt.Println("Steps:", steps)

	steps = steps + 2000
	fmt.Println("Steps:", steps)
	fmt.Println("Хорошая работа! Вы уже на пути к своей ежедневной цели!")

	// task 4
	fmt.Println()
	fmt.Println("------ Task 4 ---")

	var largeNumber int
	largeNumber = 5000000
	fmt.Println("Large Number:", largeNumber)

	// task 5
	fmt.Println()
	fmt.Println("------ Task 5 ---")

	const breakTime = 15
	fmt.Println("Break Time:", breakTime)

	// breakTime = 20
	// Мы не можем изменить значение константы
}
