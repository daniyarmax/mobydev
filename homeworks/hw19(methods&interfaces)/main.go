package main

import (
	"fmt"
	"strings"
)

// task 1
type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) area() float64      { return r.height * r.width }
func (r Rectangle) perimeter() float64 { return 2 * (r.height + r.width) }

type Shape interface {
	area() float64
	perimeter() float64
}

// task 2
type Student struct {
	name   string
	grades []float64
}

func (s Student) average() float64 {
	var sum float64 = 0.0

	for _, v := range s.grades {
		sum += v
	}

	return sum / float64((len(s.grades)))
}

type Person interface {
	average() float64
}

// task 3
type BankAccount struct {
	owner   string
	balance float64
}

func (b *BankAccount) deposit(amount float64) {
	b.balance += amount
}

func (b *BankAccount) withdraw(amount float64) {
	if b.balance >= amount {
		b.balance -= amount
	} else {
		fmt.Println("Недостотачно средств.")
	}
}

type Account interface {
	deposit(amount float64)
	withdraw(amount float64)
}

// task 4
type TextAnalyzer struct {
	text string
}

func (ta TextAnalyzer) countWords() int {
	if len(ta.text) == 0 {
		return 0
	}
	return (strings.Count(ta.text, " ")) + 1
}

func (ta TextAnalyzer) countSentences() int {
	if len(ta.text) == 0 {
		return 0
	} else if strings.ContainsAny(ta.text, ".!?") {
		dots := strings.Count(ta.text, ".")
		exclamations := strings.Count(ta.text, "!")
		questions := strings.Count(ta.text, "?")

		sentences := dots + exclamations + questions
		return sentences
	} else {
		return 1
	}
}

func (ta TextAnalyzer) countVowels() int {
	ta.text = strings.ToLower(ta.text)
	if strings.ContainsAny(ta.text, "aeiuoy") {
		as := strings.Count(ta.text, "a")
		es := strings.Count(ta.text, "e")
		is := strings.Count(ta.text, "i")
		us := strings.Count(ta.text, "u")
		os := strings.Count(ta.text, "o")
		ys := strings.Count(ta.text, "y")

		vowelsSum := as + es + is + us + os + ys

		return vowelsSum
	} else {
		return 0
	}
}

type Analyzer interface {
	countWords() int
	countSentences() int
	countVowels() int
}

func main() {
	// task 1
	fmt.Println("<----- TASK 1 ----->")
	r1 := Rectangle{height: 6, width: 7}
	r2 := Rectangle{height: 10, width: 10}

	shapes := []Shape{r1, r2}
	for i, v := range shapes {
		fmt.Printf("Прямоугольник r%d: площадь %f, периметр %f\n", i+1, v.area(), v.perimeter())
	}

	// task 2
	fmt.Println()
	fmt.Println("<----- TASK 2 ----->")

	var std1 Person = Student{"Dan", []float64{67, 90, 70}}
	var std2 Person = Student{"Tan", []float64{76, 80, 68}}
	var std3 Person = Student{"Kan", []float64{98, 76, 80}}

	var students = []Person{std1, std2, std3}

	for _, v := range students {
		fmt.Printf("Student: %s, average grade %.2f\n", v.(Student).name, v.average())
	}

	// task 3
	fmt.Println()
	fmt.Println("<----- TASK 3 ----->")

	var acc1 Account = &BankAccount{"dan", 0.0}
	var acc2 Account = &BankAccount{"kan", 500.0}

	acc1.deposit(5000)
	fmt.Println(acc1.(*BankAccount).balance)
	acc2.deposit(4501)
	fmt.Println(acc2.(*BankAccount).balance)

	acc1.withdraw(3999)
	acc2.withdraw(6000)

	var accounts = []Account{acc1, acc2}

	for _, v := range accounts {
		fmt.Printf("Счёт %s: баланс %.2f\n", v.(*BankAccount).owner, v.(*BankAccount).balance)
	}

	// task 4
	fmt.Println()
	fmt.Println("<----- TASK 4 ----->")

	var text Analyzer = TextAnalyzer{"If you do not walk today, you will run."}

	fmt.Printf("Слов: %d, предложений: %d, гласных: %d\n", text.(TextAnalyzer).countWords(), text.(TextAnalyzer).countSentences(), text.(TextAnalyzer).countVowels())
}
