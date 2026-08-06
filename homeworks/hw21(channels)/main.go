package main

import "fmt"

// task 1
func sendGreeting(ch chan string) {
	ch <- "Привет из горутины!"
}

// task 2
func squareWorker(n int, ch chan int) {
	ch <- n * n
}

// task 3
func emitNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

// task 4
func sumReader(ch chan int) {
	var sum int = 0
	for val := range ch {
		fmt.Println(val)
		sum += val
	}
	fmt.Println("Сумма чисел:", sum)
}

// task 5
func filterEven(input chan int, output chan int) {
	for val := range input {
		if val%2 == 0 {
			output <- val
		}
	}
	close(output)
}

// task 6
func checkChannel(ch chan int) {
	val, ok := <-ch
	if ok {
		fmt.Println("Received:", val)
	}

	_, ok = <-ch
	if !ok {
		fmt.Println("Channel is closed!")
	}
}

func main() {
	// task 1
	fmt.Println("<----- TASK 1 ----->")

	var ch chan string = make(chan string)
	go sendGreeting(ch)
	// val := <-ch
	fmt.Println(<-ch)

	// task 2
	fmt.Println()
	fmt.Println("<----- TASK 2 ----->")

	var squareChannel chan int = make(chan int)
	go squareWorker(9, squareChannel)
	squareValue := <-squareChannel
	fmt.Println(squareValue)

	// task 3
	fmt.Println()
	fmt.Println("<----- TASK 3 ----->")

	var numbersChannel chan int = make(chan int)
	go emitNumbers(numbersChannel)
	for val := range numbersChannel {
		fmt.Println(val)
	}

	// task 4
	fmt.Println()
	fmt.Println("<----- TASK 4 ----->")

	var sumChannel chan int = make(chan int, 3)
	sumChannel <- 3
	sumChannel <- 98
	sumChannel <- -33
	close(sumChannel)

	sumReader(sumChannel)

	// task 5
	fmt.Println()
	fmt.Println("<----- TASK 5 ----->")

	var inputChannel chan int = make(chan int, 10)
	var outputChannel chan int = make(chan int, 5)

	for i := 1; i <= 10; i++ {
		inputChannel <- i
	}
	close(inputChannel)
	filterEven(inputChannel, outputChannel)

	for val := range outputChannel {
		fmt.Println(val)
	}

	// task 6
	fmt.Println()
	fmt.Println("<----- TASK 6 ----->")

	var check chan int = make(chan int, 1)
	check <- 5
	close(check)

	checkChannel(check)

	// task 7
	fmt.Println()
	fmt.Println("<----- TASK 7 ----->")

	var message chan string = make(chan string, 3)
	message <- "Salem, alem"
	message <- "Hello, world"
	message <- "Privet, mir"

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)

	message <- "Assalamaleikum"
}
