package main

import "fmt"

func main() {
	const (
		BaseTariff    = 0.45 // цена за 1 кВт·ч
		HighLoadTax   = 0.15 // налог на высокое потребление
		NightDiscount = 0.30 // ночная скидка
	)

	var name string
	var power float64
	var workTime int
	var nightMode bool
	var category string

	for {
		fmt.Print("Введите название прибора (или done): ")
		fmt.Scanln(&name)

		if name == "done" {
			break
		}

		fmt.Print("Мощность (кВт): ")
		_, err := fmt.Scanln(&power)
		if err != nil {
			fmt.Println("Ошибка ввода мощности:", err)
			return
		}

		fmt.Print("Время работы (часы): ")
		_, err = fmt.Scanln(&workTime)
		if err != nil {
			fmt.Println("Ошибка ввода времени работы:", err)
			return
		}

		fmt.Print("Ночной режим (true/false): ")
		_, err = fmt.Scanln(&nightMode)
		if err != nil {
			fmt.Println("Ошибка ввода ночного режима:", err)
			return
		}

		// расчёт
		consumption := (power * float64(workTime)) / 1000 // кВт·ч
		total := consumption * BaseTariff

		// ночная скидка
		if nightMode {
			total -= total * NightDiscount
		}

		// если суммарное потребление за работу >100 кВт·ч, применяем налог
		if power > 10 {
			total += total * HighLoadTax
		}

		switch {
		case power < 100:
			category = "Экономный"
		case power > 1000:
			category = "Мощный"
		default:
			category = "Стандартный"
		}

		fmt.Println("--- Отчёт по прибору ---")
		fmt.Printf("Прибор: %s [Категория: %s]\nПотребление: %.2f кВт·ч\nИтоговая стоимость: %.2f\n", name, category, consumption, total)
	}
}
