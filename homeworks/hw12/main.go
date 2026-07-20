package main

import "fmt"

func main() {

	// task 1
	fmt.Println("<----- Task 1 ----->")
	var runningExercises = [...]string{"runningexercise1", "runningexercise2", "runningexercise3"}
	var walkingExercise = [...]string{"walkingexercise1", "walkingexercise2", "walkingexercise3"}

	fmt.Println("number of running exercise: ", len(runningExercises))
	fmt.Println("number of walking exercise: ", len(walkingExercise))

	// task 2
	fmt.Println()
	fmt.Println("<----- Task 2 ----->")

	var subjectList = [3]string{"Physics", "Chemistry", "Geography"}

	firstElement := subjectList[0]
	lastElement := subjectList[len(subjectList)-1]

	fmt.Println("Subjects: ", subjectList)
	fmt.Println("First subject:", firstElement)
	fmt.Println("Last subject:", lastElement)

	subjectList[1] = "Biology"

	fmt.Println("Changed subject list:", subjectList)

	// task 3
	fmt.Println()
	fmt.Println("<----- Task 3 ----->")

	var data = [3]string{"Tom", "35", "New York"}
	name := data[0]
	age := data[1]
	home := data[2]

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Hometown:", home)

	// task 4
	fmt.Println()
	fmt.Println("<----- Task 4 ----->")

	var numberList = [5]int{1, 2, 3, 4, 5}

	threeExists := false
	for _, val := range numberList {
		if val == 3 {
			fmt.Println("Число 3 есть в массиве")
			threeExists = true
			break
		}
	}

	if !threeExists {
		fmt.Println("Число 3 отсутствует в массиве")
	}

	// task 5
	fmt.Println()
	fmt.Println("<----- Task 5 ----->")

	var friendList = [3]string{"Aibek", "Nurdaulet", "Orken"}

	AibekExists := false
	for _, val := range friendList {
		if val == "Aibek" {
			fmt.Println("Мне повезло")
			AibekExists = true
			break
		}
	}

	if !AibekExists {
		fmt.Println("Мне не повезло")
	}

	// task 6
	fmt.Println()
	fmt.Println("<----- Task 6 ----->")

	var firstList = [3]int{1, 2, 3}
	var secondList = [3]int{1, 2, 3}

	if firstList == secondList {
		fmt.Println("Lists are equal")
	} else {
		fmt.Println("Lists are not equal")
	}

	// task 7
	fmt.Println()
	fmt.Println("<----- Task 7 ----->")

	var myWishList = [3]string{"Airpods", "Book", "Watch"}
	var friendsWishList = [2]string{"Phone", "Coke"}

	var registrationList [5]string

	currentIndex := 0

	for _, item := range myWishList {
		registrationList[currentIndex] = item
		currentIndex++
	}

	for _, item := range friendsWishList {
		registrationList[currentIndex] = item
		currentIndex++
	}

	fmt.Println("All wishlist:", registrationList)

	// task 8
	fmt.Println()
	fmt.Println("<----- Task 8 ----->")

	var toyList = [3]string{"Car", "Doll", "Ball"}
	var testToyList = toyList

	testToyList[1] = "Boat"

	fmt.Println("Toy list:", toyList)
	fmt.Println("Test toy list:", testToyList)
}
