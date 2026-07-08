package main

import (
	"fmt"
)

func main() {
	const (
		baseRate     = 5.50
		taxRate      = 0.12
		distanceRate = 2.0
		fragileFee   = 0.2
	)

	var name string
	var distance float64
	var weight float64
	var numberOfFragileItems int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter the distance (in km): ")
	fmt.Scanln(&distance)

	fmt.Print("Enter the weight of the package (in kg): ")
	fmt.Scanln(&weight)

	fmt.Print("Enter the number of fragile items: ")
	fmt.Scanln(&numberOfFragileItems)

	baseCost := (weight*baseRate)*(1+fragileFee*float64(numberOfFragileItems)) + (distance * distanceRate)
	finalCost := baseCost * (1 + taxRate)
	fmt.Println("--- Delivery information ---")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Final cost: %.2f\n", finalCost)
}
