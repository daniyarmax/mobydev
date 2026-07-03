package main

import (
	"fmt"
	"math"
)

func main() {
	// task 1
	fmt.Println("----- Task 1 -----")

	var bannerWidth int = 12
	var bannerHeight int = 8
	var bannerArea int = bannerWidth * bannerHeight
	fmt.Println("Banner are:", bannerArea)

	halfBannerAred := bannerArea / 2
	fmt.Println("Half banner area:", halfBannerAred)

	bannerBorderLength := 2 * (bannerWidth + bannerHeight)
	fmt.Println("Banner border length:", bannerBorderLength)

	// task 2
	fmt.Println()
	fmt.Println("----- Task 2 -----")

	boxCount := 29
	leftoverBoxes := boxCount % 5
	fmt.Println("Leftover boxes:", leftoverBoxes)

	// task 3
	fmt.Println()
	fmt.Println("----- Task 3 -----")

	tempMorning := 16
	tempAfternoon := 23
	tempEvening := 18

	totalTemp := tempMorning + tempAfternoon + tempEvening
	avgTemp := totalTemp / 3
	fmt.Println("Average temperature:", avgTemp)

	// task 4
	fmt.Println()
	fmt.Println("----- Task 4 -----")

	knownWords := 67
	goalWords := 120

	progressPercent := float64(knownWords) / float64(goalWords) * 100
	fmt.Printf("Progress: %.2f%%\n", progressPercent)

	// task 5
	fmt.Println()
	fmt.Println("----- Task 5 -----")

	coins := 0

	coins += 500
	fmt.Println("Coins after collecting 500:", coins)

	coins += 1200
	fmt.Println("Coins after collecting 1200:", coins)

	coins /= 2
	fmt.Println("Coins after spending half:", coins)

	coins -= 300
	fmt.Println("Coins after spending 300:", coins)

	// task 6
	fmt.Println()
	fmt.Println("----- Task 6 -----")

	participants := 42
	groups := 8

	participantsPerGroup := participants / groups

	fmt.Println("Participants per group:", participantsPerGroup)

	// task 7
	fmt.Println()
	fmt.Println("----- Task 7 -----")

	firstExpression := 20 - 4*3
	secondExpression := (20 - 4) * 3

	fmt.Println("First expression result:", firstExpression)
	fmt.Println("Second expression result:", secondExpression)
	// Результаты различаются из-за порядка выполнения операций.
	// В первом выражении сначала выполняется умножение (4*3=12), затем вычитание (20-12=8).
	// Во втором выражении сначала выполняется вычитание в скобках
	// (20-4=16), затем умножение (16*3=48).

	// task 8
	fmt.Println()
	fmt.Println("----- Task 8 -----")

	squareValue := 81
	squareRoot := math.Sqrt(float64(squareValue))
	fmt.Println("Square root of", squareValue, "is", squareRoot)

	multiplier := 5
	exponent := 2
	result := math.Pow(float64(multiplier), float64(exponent))
	fmt.Println("Result of", multiplier, "raised to the power of", exponent, "is", result)

}
