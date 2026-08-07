package main

import (
	"errors"
	"fmt"
	"math"
)

// task 2
func divide(a, b int) (int, interface{}) {
	if b == 0 {
		return 0, "Ошибка: деление на ноль невозможно"
	}
	return a / b, nil
}

// task 3
func checkAge(age int) error {
	if age > 120 || age < 0 {
		return errors.New("указан некорректный возраст")
	} else {
		return nil
	}
}

// task 4
func validatePassword(password string) (bool, error) {
	if len(password) < 6 {
		return false, errors.New("пароль слишком короткий")
	} else {
		return true, nil
	}
}

// task 5
type error interface {
	Error() string
}
type invalidRadiusError struct{}

func (err invalidRadiusError) Error() string {
	return "радиус не может быть отрицательным"
}

// task 6
func calculateCircleArea(radius int) (float64, error) {
	if radius < 0 {
		return 0, invalidRadiusError{}
	}
	return math.Pi * float64(radius) * float64(radius), nil
}

// task 7
func findUser(users []string, user string) (int, error) {
	for i, v := range users {
		if v == user {
			return i, nil
		}
	}
	return -1, errors.New("пользователь не найден")
}

func main() {
	// task 1
	fmt.Println("<----- TASK 1 ----->")

	var goods = map[string]int{
		"Book":     1500,
		"Notebook": 860,
		"Pen":      50,
		"Pencil":   50,
	}

	good := "Pen"

	val, ok := goods[good]
	if ok {
		fmt.Printf("Цена товара: %d тг\n", val)
	} else {
		fmt.Println("Товар отсутсвует!")
	}

	// task 2
	fmt.Println()
	fmt.Println("<----- TASK 2 ----->")

	result, err := divide(12, 0)
	if err != nil {
		fmt.Println("Ошибка")
	} else {
		fmt.Println("Результат деление: ", result)
	}

	// task 3
	fmt.Println()
	fmt.Println("<----- TASK 3 ----->")

	fmt.Println(checkAge(3))
	fmt.Println(checkAge(-1))
	fmt.Println(checkAge(121))

	// task 4
	fmt.Println()
	fmt.Println("<----- TASK 4 ----->")

	password := "alksfjdlkj"
	fmt.Println(validatePassword(password))

	password = "12234"
	fmt.Println(validatePassword(password))

	// task 5
	fmt.Println()
	fmt.Println("<----- TASK 5 ----->")

	invRadErr := invalidRadiusError{}
	fmt.Println(invRadErr.Error())

	// task 6
	fmt.Println()
	fmt.Println("<----- TASK 6 ----->")

	fmt.Println(calculateCircleArea(5))
	fmt.Println(calculateCircleArea(0))

	// task 7
	fmt.Println()
	fmt.Println("<----- TASK 7 ----->")

	users := []string{"Daniyar", "Daniyar1"}
	fmt.Println(findUser(users, "daniyar"))
	fmt.Println(findUser(users, "Daniyar"))

}
