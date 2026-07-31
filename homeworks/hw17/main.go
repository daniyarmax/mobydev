package main

import "fmt"

// task 1
func square(number int) int {
	return number * number
}

// task 2
func maxNumber(num1, num2 int) int {
	var maxNumber int
	if num1 > num2 {
		maxNumber = num1
	} else {
		maxNumber = num2
	}

	return maxNumber
}

// task 3
func isEven(num1 int) bool {
	if num1%2 == 0 {
		return true
	} else {
		return false
	}
}

// task 4
func greetUser(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// task 5
func sumSlice(slice1 []int) int {
	sum := 0

	for _, val := range slice1 {
		sum += val
	}

	return sum
}

// task 6
func checkLogin(login, password string) bool {
	data := map[string]string{"admin": "1234"}

	if data[login] == password {
		return true
	} else {
		return false
	}
}

// task 7
func increaseBalance(balance *int, topUp int) {
	*balance += topUp
}

// task 8
func resetAttempts(attempts *int) {
	if *attempts > 3 {
		*attempts = 0
	}
}

//------------------------------->

func main() {
	// task 1
	fmt.Println("<----- TASK 1 ----->")

	fmt.Println(square(5))
	fmt.Println(square(6))
	fmt.Println(square(11))

	// task 2
	fmt.Println()
	fmt.Println("<----- TASK 2 ----->")

	fmt.Println(maxNumber(67, 911))
	fmt.Println(maxNumber(666, 667))

	// task 3
	fmt.Println()
	fmt.Println("<----- TASK 3 ----->")

	fmt.Println(isEven(888))
	fmt.Println(isEven(-777))
	fmt.Println(isEven(0))

	// task 4
	fmt.Println()
	fmt.Println("<----- TASK 4 ----->")

	name := "Danchik"
	greetUser(name)

	// task 5
	fmt.Println()
	fmt.Println("<----- TASK 5 ----->")

	slice1 := []int{1, 3, 67, 666, 667}
	fmt.Println("Sum of slice:", sumSlice(slice1))

	// task 6
	fmt.Println()
	fmt.Println("<----- TASK 6 ----->")

	login := "admin"
	password := "1234"

	if checkLogin(login, password) {
		fmt.Println("Credentials are correct.")
	} else {
		fmt.Println("Credentials are not correct.")
	}

	// task 7
	fmt.Println()
	fmt.Println("<----- TASK 7 ----->")

	balance := 5000
	topUp := 999

	increaseBalance(&balance, topUp)
	fmt.Println(balance)

	// task 8
	fmt.Println()
	fmt.Println("<----- TASK 8 ----->")

	attempts := 4
	resetAttempts(&attempts)
	fmt.Println(attempts)
}
