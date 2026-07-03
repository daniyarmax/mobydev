package main

import (
	"fmt"
)

func main() {
	// task 1
	fmt.Println("----- Task 1 -----")
	var age int = 19
	fmt.Println("Age:", age)

	age += 1
	fmt.Println("New Age:", age)

	// task 2
	fmt.Println()
	fmt.Println("----- Task 2 -----")

	var height int = 170
	fmt.Println("Height in cm: ", height)

	var heightInMeters float64
	heightInMeters = float64(height) / 100.0
	fmt.Println("Height in meters: ", heightInMeters)

	// task 3
	fmt.Println()
	fmt.Println("----- Task 3 -----")

	var isStudent bool = true
	fmt.Println("Is student:", isStudent)

	// task 4
	fmt.Println()
	fmt.Println("----- Task 4 -----")

	var temperature float64 = 6.7
	fmt.Println("Temperature outside:", temperature)
	fmt.Println("Outside is cold")

	// task 5
	fmt.Println()
	fmt.Println("----- Task 5 -----")

	var favoriteQuote string = "The only limit to our realization of tomorrow is our doubts of today."
	fmt.Println("Favorite quote:", favoriteQuote)

	// task 6
	fmt.Println()
	fmt.Println("----- Task 6 -----")

	var pi float64
	pi = 3.14
	fmt.Println("Value of pi:", pi)

	//pi = "3.14159"
	// Эта строка вызовет ошибку компиляции, так как переменная pi объявлена как float64, а мы пытаемся присвоить ей строку. В Go нельзя изменять тип переменной после её объявления.
}
