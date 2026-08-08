package main

import "fmt"

// task 1
func printData() {
	fmt.Println("Обработка программы")
}

// task 3
func checkAge(age int) {
	if age < 0 {
		panic("Age can't be negative!")
	}
}

// task 4
func cleanup() {
	fmt.Println("Очистка ресурсов выполнена")
}

// task 5
func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Error:", r)
	}
}

// task 6
func safeDivide(a, b float64) float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error: ", r)
		}
	}()

	if b == 0 {
		panic("Dividing by zero!")
	}
	return a / b
}

// task 7
func getElement(slice []int, index int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	if index >= len(slice) {
		panic("Out of bound")
	}
	return slice[index]
}

func main() {
	// task 1
	fmt.Println("<----- TASK 1 ----->")

	fmt.Println("Старт программы.")
	// defer printData()
	fmt.Println("Завершение программы.")

	// task 2
	fmt.Println()
	fmt.Println("<----- TASK 2 ----->")

	// defer fmt.Println("First")
	// defer fmt.Println("Second")
	// defer fmt.Println("Third")

	// task 3
	fmt.Println()
	fmt.Println("<----- TASK 3 ----->")

	// checkAge(-1)

	// task 4
	fmt.Println()
	fmt.Println("<----- TASK 4 ----->")

	// defer cleanup()
	// panic("Критическая ошибка!")

	// task 5
	fmt.Println()
	fmt.Println("<----- TASK 5 ----->")

	// defer handlePanic()
	// panic("Errorski")

	// task 6
	fmt.Println()
	fmt.Println("<----- TASK 6 ----->")
	safeDivide(1, 0)

	// task 7
	fmt.Println()
	fmt.Println("<----- TASK 7 ----->")

	slice := []int{1, 2, 3, 4}
	fmt.Println(getElement(slice, 10))
}
