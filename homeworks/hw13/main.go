package main

import "fmt"

func main() {
	// task 1
	fmt.Println("<----- Task 1 ----->")

	var numbers [3][5]int

	var maxNumber int = numbers[0][0]
	var minColumn int = 0
	var row int
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[i]); j++ {
			var number int
			fmt.Println("Enter a number to fill the matrix:")
			fmt.Scanln(&number)

			numbers[i][j] = number

			if number > maxNumber {
				maxNumber = number
				row = i
				minColumn = j
			} else if number == maxNumber {
				if minColumn > j {
					minColumn = j
				}
			}
		}
	}

	fmt.Println(numbers)
	fmt.Println("Indexes of the maximum element:", row, minColumn)

	// task 2
	fmt.Println()
	fmt.Println("<----- Task 2 ----->")

	var uk [11][11]string

	for i := 0; i < len(uk); i++ {
		for j := 0; j < len(uk); j++ {
			if (i == j) || (i+j == len(uk)-1) || (i == 5) || (j == 5) {
				uk[i][j] = "*"
			} else {
				uk[i][j] = "."
			}
		}
		fmt.Println(uk[i])
	}

	// task 3
	fmt.Println()
	fmt.Println("<----- Task 3 ----->")

	var chessBoard [8][8]string

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				chessBoard[i][j] = "."
			} else {
				chessBoard[i][j] = "*"
			}
		}
		fmt.Println(chessBoard[i])
	}

	// task 4
	fmt.Println()
	fmt.Println("<----- Task 4 ----->")

	var numbers1 [4][4]int
	for i := 0; i < len(numbers1); i++ {
		for j := 0; j < len(numbers1[i]); j++ {
			var number1 int
			fmt.Println("Enter a number to fill the matrix:")
			fmt.Scanln(&number1)

			numbers1[i][j] = number1
		}

	}
	fmt.Println(numbers1)
	var i int
	var j int
	var temp [4]int

	fmt.Println("Enter value i:")
	fmt.Scanln(&i)
	fmt.Println("Enter value j:")
	fmt.Scanln(&j)

	temp = numbers1[i]
	numbers1[i] = numbers1[j]
	numbers1[j] = temp

	fmt.Println(numbers1)

	// task 5
	fmt.Println()
	fmt.Println("<----- Task 5 ----->")

	var numbers2 [4][4]int
	for i := 0; i < len(numbers2); i++ {
		for j := 0; j < len(numbers2[i]); j++ {
			var number2 int
			fmt.Println("Enter a number to fill the matrix:")
			fmt.Scanln(&number2)

			numbers2[i][j] = number2
		}

	}
	fmt.Println(numbers2)
	var i1 int
	var j1 int

	fmt.Println("Enter value i:")
	fmt.Scanln(&i1)
	fmt.Println("Enter value j:")
	fmt.Scanln(&j1)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j == i1 {
				temp1 := numbers2[i][i1]
				numbers2[i][i1] = numbers2[i][j1]
				numbers2[i][j1] = temp1
			} else {
				break
			}
		}
	}

	fmt.Println(numbers2)
}
