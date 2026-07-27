package main

import (
	"fmt"
	"reflect"
	"slices"
)

func main() {
	// task 1
	fmt.Println("<----- task 1 ----->")

	numbers := []int{}
	sum := 0

	for true {
		var number int
		fmt.Print("Enter a number (or 0 to stop): ")
		fmt.Scanln(&number)
		if number == 0 {
			break
		}

		sum += number
		numbers = append(numbers, number)
	}

	fmt.Println("Slice: ", numbers)
	fmt.Println("Sum of slice: ", sum)

	// task 2
	fmt.Println()
	fmt.Println("<----- task 2 ----->")

	values := []int{}
	evenNumbers := []int{}

	for true {
		var value int
		fmt.Print("Enter a number (or 0 to stop): ")
		fmt.Scanln(&value)
		if value == 0 {
			break
		} else if value%2 == 0 {
			evenNumbers = append(evenNumbers, value)
		}

		values = append(values, value)
	}

	fmt.Println("Numbers: ", values)
	fmt.Println("Even numbers: ", evenNumbers)

	// task 3
	fmt.Println()
	fmt.Println("<----- task 3 ----->")

	data := []int{}

	for true {
		var num int
		fmt.Print("Enter a number (or 0 to stop): ")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}

		data = append(data, num)
	}

	if len(data) >= 3 {
		i := 2
		data = append(data[:i], data[i+1:]...)
	}

	fmt.Println("Data: ", data)

	// task 4
	fmt.Println()
	fmt.Println("<----- task 4 ----->")

	temps := []int{}
	maxTemp := -10000
	minTemp := 10000
	for true {
		var temp int
		fmt.Print("Enter a temperature value (or 0 to stop): ")
		fmt.Scanln(&temp)
		if temp == 0 {
			break
		} else if maxTemp < temp {
			maxTemp = temp
		} else if minTemp > temp {
			minTemp = temp
		}

		temps = append(temps, temp)
	}
	fmt.Println("Maximum temperature: ", maxTemp)
	fmt.Println("Minimum temperature: ", minTemp)

	// task 5
	fmt.Println()
	fmt.Println("<----- task 5 ----->")

	words := []string{}

	for true {
		var word string
		fmt.Print("Enter a word (or write 'stop' to stop): ")
		fmt.Scanln(&word)
		if word == "stop" {
			break
		}

		words = append(words, word)
	}

	reversed := make([]string, len(words))

	for i, word := range words {
		reversed[len(words)-1-i] = word
	}

	fmt.Println("Basic slice:", words)
	fmt.Println("Reversed slice:", reversed)

	// task 6
	fmt.Println()
	fmt.Println("<----- task 6 ----->")

	nums := []int{}
	for true {
		var num int
		fmt.Print("Enter a number (or 0 to stop): ")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}

		nums = append(nums, num)
	}

	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	slices.Sort(sortedNums)

	fmt.Println("Numbers: ", nums)
	fmt.Println("Sorted numbers: ", sortedNums)
	fmt.Println(reflect.DeepEqual(sortedNums, nums))
}
