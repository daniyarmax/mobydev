package main

import "fmt"

func main() {
	// task 1
	fmt.Println("----- Task 1-----")

	var i int = 1
	for ; i <= 20; i++ {
		fmt.Println(i)
	}

	// task 2
	fmt.Println()
	fmt.Println("----- Task 2-----")

	sum := 0

	for i := 1; i <= 100; i++ {
		sum += i

	}

	fmt.Println("Сумма чисел от 1 до 100:", sum)

	// task 3
	fmt.Println()
	fmt.Println("----- Task 3-----")

	var number int

	fmt.Print("Введите число для вывода таблицы умножения: ")
	fmt.Scanln(&number)

	for i := 1; i <= 10; i++ {
		fmt.Println(number, " * ", i, " = ", number*i)
	}

	// task 4
	fmt.Println()
	fmt.Println("----- Task 4-----")

	var n int

	fmt.Print("Введите число для проверки кратности 3: ")
	fmt.Scanln(&n)

	for i := 3; i <= n; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	// task 5
	fmt.Println()
	fmt.Println("----- Task 5-----")

	var number1 int
	fmt.Print("Введите число для подсчета цифр: ")
	fmt.Scanln(&number1)

	digitCount := 1

	for number1 >= 10 {
		number1 /= 10
		digitCount++
	}
	fmt.Println("Количество цифр в числе:", digitCount)

	// task 6
	fmt.Println()
	fmt.Println("----- Task 6-----")

	var text string

	fmt.Print("Введите текст: ")
	fmt.Scanln(&text)

	for _, char := range text {
		fmt.Printf("%c\n", char)
	}

	// task 7
	fmt.Println()
	fmt.Println("----- Task 7-----")

	balance := 3000

	for {
		var command int
		fmt.Println("Введите команду: ")
		fmt.Scanln(&command)
		switch command {
		case 1:
			fmt.Println("Баланс:", balance)
		case 2:
			balance += 500
			fmt.Println("+500 к балансу. Баланс:", balance)
		case 3:
			balance -= 200
			fmt.Println("-200 с баланса. Баланс:", balance)

		case 0:
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неверная команда. Попробуйте снова.")
		}
	}

}
