package main

import "fmt"

func main() {
	// task 1
	fmt.Println("----- Task 1 -----")
	var text string = "Just do it!"
	const author string = "Daniyar"

	fmt.Printf("Автор %s написал: %s\n", author, text)
	fmt.Println("Длина строки: ", len(text))

	// task 2
	fmt.Println()
	fmt.Println("----- Task 2 -----")

	var message string = "\"Salem, alem!\""

	if message == "" {
		fmt.Println("Пустая строка!")
	} else {
		fmt.Println("Строка не пустая")
		fmt.Println(message)
	}

	// task 3
	fmt.Println()
	fmt.Println("<----- Task 3 ----->")

	var input string
	fmt.Scan(&input)

	if len(input) < 5 {
		fmt.Println("Слишком короткая слово")
	} else if len(input) >= 5 && len(input) <= 10 {
		fmt.Println("Нормальная длина")
	} else {
		fmt.Println("Слишком длинное слово")
	}

	// task 4
	fmt.Println()
	fmt.Println("<----- Task 4 ----->")

	var word string
	word = "Amanbol"
	symbol1 := word[4]
	fmt.Printf("%c", symbol1)
}
