package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Эспрессо": 800,
		"Латте":    1200,
		"Капучино": 1100,
		"Сэндвич":  1500,
		"Круассан": 900,
	}

	order := []string{} // слайс для хранение данных
	costBeforeDiscount := 0.0
	for true {
		var value string
		fmt.Println()
		fmt.Println("Выберите блюдо:")
		for key, val := range menu {
			fmt.Printf(" %-10s -> %f\n", key, val)
		}
		fmt.Print("Название блюдо: ")
		fmt.Scanln(&value)
		if value == "exit" {
			break
		} else {
			_, ok := menu[value]
			if ok {
				order = append(order, value)
				fmt.Println("Блюдо добавлен в корзину.")
				fmt.Print("Корзина: ")
				fmt.Println(order)
				costBeforeDiscount += menu[value]
			} else {
				fmt.Println("К сожалению, этого блюда нет в меню.")
			}
		}

	}

	total := 0.0
	if costBeforeDiscount >= 5000 {
		total = costBeforeDiscount * 0.9
	} else {
		total = costBeforeDiscount
	}
	fmt.Println("Ваша корзина: ")
	for idx, val := range order {
		fmt.Printf("%d.  %-10s -> %f тг\n", idx+1, val, menu[val])
	}

	fmt.Printf("Сумма без скидки: %f тг\n", costBeforeDiscount)
	fmt.Printf("Скидка: %f тг\n", costBeforeDiscount-total)
	fmt.Printf("ИТОГО: %f тг\n", total)
}
