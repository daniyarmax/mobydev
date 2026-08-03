package main

import (
	"fmt"
	"sync"
	"time"
)

// task 1
func printInfo() {
	fmt.Println("Goroutine has started.")
}

// task 2
func sayHello(name string) {
	fmt.Println("Hello,", name)
}

// task 3
func heavyTask(wg *sync.WaitGroup) {
	fmt.Println("Выполняю задачу...")
	wg.Done()
}

// task 4
func count(id int, wg *sync.WaitGroup) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Goroutine %d: number %d\n", id, i)
	}

	wg.Done()
}

// task 5
func calculateSquare(num int, wg *sync.WaitGroup) {
	fmt.Println(num * num)
	wg.Done()
}

// task 6

func checkStatus(site string, wg *sync.WaitGroup) {
	fmt.Printf("Сайт %s доступен\n", site)
	wg.Done()
}

// task 7
func processData(id int, category string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Запуск процесса %d в категории %s\n", id, category)
	time.Sleep(time.Duration(delay.Seconds()))

	fmt.Printf("Процесс %d завершен.\n", id)

}

func main() {
	// task 1
	fmt.Println("<----- TASK 1 ----->")
	go printInfo()

	time.Sleep(time.Second)

	// task 2
	fmt.Println()
	fmt.Println("<----- TASK 2 ----->")

	go sayHello("Daniyar")
	go sayHello("Mama")
	go sayHello("Bratan")

	time.Sleep(time.Second)

	// task 3
	fmt.Println()
	fmt.Println("<----- TASK 3 ----->")

	var wg sync.WaitGroup

	wg.Add(1)
	go heavyTask(&wg)

	wg.Wait()

	// task 4
	fmt.Println()
	fmt.Println("<----- TASK 4 ----->")

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go count(i, &wg)

	}
	wg.Wait()

	// task 5
	fmt.Println()
	fmt.Println("<----- TASK 5 ----->")

	var numbers = []int{2, 4, 73, 65, 11}
	for _, v := range numbers {
		wg.Add(1)
		go calculateSquare(v, &wg)
	}

	wg.Wait()

	// task 6
	fmt.Println()
	fmt.Println("<----- TASK 6 ----->")

	var sites = []string{"google.com", "github.com", "instagram.com"}

	for _, v := range sites {
		wg.Add(1)
		go checkStatus(v, &wg)
	}

	wg.Wait()

	// task 7
	for i := 0; i < 3; i++ {
		wg.Add(1)
		var category string
		var delay time.Duration

		fmt.Print("Enter a category: ")
		fmt.Scanln(&category)
		fmt.Print("Enter a delay: ")
		fmt.Scanln(&delay)

		go processData(i+1, category, delay, &wg)
		wg.Wait()
	}

}
