package main

import (
	"fmt"
)

func main() {
	fmt.Println("----- Task 1 -----")

	// true
	fmt.Println(5 == 5)

	// true
	fmt.Println(10 != 3)

	// false
	fmt.Println(7 > 12)

	// true
	fmt.Println(15 < 20)

	// true
	fmt.Println(8 >= 8)

	// false
	fmt.Println(6 <= 4)

	// false
	fmt.Println((10 > 5) && (3 < 1))

	// true
	fmt.Println((10 > 5) || (3 < 1))

	// false
	fmt.Println(!(5 == 5))

	// true
	fmt.Println(!(7 < 3))

	// false
	fmt.Println(true && false)

	// false
	fmt.Println(false || false)

	// true
	fmt.Println(true || false)

	// true
	fmt.Println((4+6 == 10) && (9 > 2))

	// true
	fmt.Println((12/3 == 4) || (8 < 5))

	// task 2
	fmt.Println()
	fmt.Println("----- Task 2 -----")

	var age int = 20
	var hasTicket bool = true

	canEnter := age >= 18 && hasTicket

	fmt.Println("Age:", age)
	fmt.Println("Has ticket:", hasTicket)
	fmt.Println("Can enter:", canEnter)

	// task 3
	fmt.Println()
	fmt.Println("----- Task 3 -----")

	var isLoggedIn bool = true
	var isAdmin bool = true

	hasAccess := (isLoggedIn && isAdmin) || (isLoggedIn && !isAdmin)

	fmt.Println("Is logged in:", isLoggedIn)
	fmt.Println("Is admin:", isAdmin)
	fmt.Println("Has access:", hasAccess)
}
