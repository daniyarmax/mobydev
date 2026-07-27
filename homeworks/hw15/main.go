package main

import (
	"fmt"
	"reflect"
)

func main() {
	// task 1
	fmt.Println("<----- TASK 1 ----->")

	var toolUsage map[string]int

	toolUsage = map[string]int{
		"Golang": 3,
		"VSCode": 1,
		"GitHub": 5,
	}

	for key, ok := range toolUsage {
		fmt.Printf("Tool: %s, Usage count: %d\n", key, ok)
	}

	// task 2
	fmt.Println()
	fmt.Println("<----- TASK 2 ----->")

	var buildStatus map[string]bool

	buildStatus = map[string]bool{
		"build": true,
		"run":   false,
	}

	if buildStatus["build"] {
		fmt.Println("Сборка прошла успешно")
	}

	// task 3
	fmt.Println()
	fmt.Println("<----- TASK 3 ----->")

	var name string
	fmt.Print("Введите имя: ")
	fmt.Scanln(&name)

	var userInfo = map[string]string{
		"name":       name,
		"isLoggedIn": "true",
	}

	fmt.Printf("Пользователь %s вошёл в систему", userInfo["name"])

	// task 4
	fmt.Println()
	fmt.Println("<----- TASK 4 ----->")

	var cpuLoad = map[int]int{
		1: 40,
		2: 65,
		3: 30,
		4: 14,
	}

	maxValue := 0
	maxKey := 0
	for key, val := range cpuLoad {
		if val > maxValue {
			maxValue = val
			maxKey = key
		}
	}

	fmt.Println("Ядро с самым максимальным загрузкой: ", maxKey)

	// task 5
	fmt.Println()
	fmt.Println("<----- TASK 5 ----->")

	var examResults = map[string]int{
		"Pinkman":      67,
		"Walter White": 99,
		"Gustavo":      95,
		"Goodman":      80,
	}

	for key, val := range examResults {
		if val >= 70 && key == "Walter White" {
			fmt.Printf("%s будет получать стипендию. Оценка: %d, Yeah Mr.White, yeah science!\n", key, val)
		} else if val >= 70 {
			fmt.Printf("%s будет получать стипендию. Оценка: %d\n", key, val)
		} else {
			fmt.Printf("%s НЕ будет получать стипендию. Оценка: %d\n", key, val)
		}
	}

	// task 6
	fmt.Println()
	fmt.Println("<----- TASK 6 ----->")

	var words = []string{"go", "is", "fast"}
	var wordLength = make(map[string]int)

	for _, val := range words {
		wordLength[val] = len(val)
	}

	fmt.Println(wordLength)

	// task 7
	fmt.Println()
	fmt.Println("<----- TASK 7 ----->")

	menu := map[string]int{
		"Burger": 2500,
		"Pizza":  1100,
		"Doner":  1790,
	}

	var mealName string
	fmt.Print("Enter a meal name: ")
	fmt.Scanln(&mealName)

	val, okay := menu[mealName]
	if okay {
		fmt.Println("Price: ", val)
	} else {
		fmt.Println("Блюдо не найдено")
	}

	// task 8
	fmt.Println()
	fmt.Println("<----- TASK 8 ----->")

	loginAttempts := map[string]int{
		"admin": 2,
		"guest": 0,
	}

	loginAttempts["admin"] += 1

	if loginAttempts["admin"] > 2 {
		fmt.Println("Доступ заблокирован")
	}

	// task 9
	fmt.Println()
	fmt.Println("<----- TASK 9 ----->")

	sales := [][]int{
		{2, 4, 6},
		{1, 3, 5},
	}

	total := make(map[int]int)

	for i := 0; i < len(sales); i++ {
		total[i+1] = 0
		for j := 0; j < len(sales[i]); j++ {
			total[i+1] += sales[i][j]
		}
	}

	fmt.Println(total)

	// task 10
	fmt.Println()
	fmt.Println("<----- TASK 10 ----->")

	numbers := []int{4, 7, 9, 2, 5}
	numberStatus := make(map[int]string)

	for _, val := range numbers {
		if val%2 == 0 {
			numberStatus[val] = "even"
		} else {
			numberStatus[val] = "odd"
		}
	}

	fmt.Println(numberStatus)

	// task 11
	fmt.Println()
	fmt.Println("<----- TASK 11 ----->")

	defaultConfig := map[string]string{
		"host": "localhost",
		"port": "8080",
		"mode": "production",
	}

	currentConfig := map[string]string{
		"host": "localhost",
		"port": "8080",
		"mode": "production",
	}

	if reflect.DeepEqual(defaultConfig, currentConfig) {
		fmt.Println("Конфигурации совпадают")
	} else {
		fmt.Println("Конфигурации отличаются")
	}

	currentConfig["mode"] = "debug"

	if reflect.DeepEqual(defaultConfig, currentConfig) {
		fmt.Println("Конфигурации совпадают")
	} else {
		fmt.Println("Конфигурации отличаются")
	}
}
