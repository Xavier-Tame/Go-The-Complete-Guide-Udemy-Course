package main

import (
	"fmt"
	"os"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Enter total revenue: ")
	// fmt.Scan(&revenue)

	revenue := getUserInput("Enter total revenue: ")

	// fmt.Print("Enter total expenses: ")
	// fmt.Scan(&expenses)

	expenses := getUserInput("Enter total expenses: ")

	// fmt.Print("Enter tax rate as whole number (e.g., 20 for 20%): ")
	// fmt.Scan(&taxRate)

	taxRate := getUserInput("Enter tax rate as whole number (e.g., 20 for 20%): ")

	if revenue <= 0 || expenses <= 0 || taxRate <= 0 {
		fmt.Println("Error: Inputs must be non-negative numbers and greater than zero.")
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Net Income: %.1f\n", profit)
	fmt.Printf("EBT to Net Income Ratio: %.2f\n", ratio)

	saveFinancialsToFile(ebt, profit, ratio)
}

func getUserInput(text string) float64 {
	var userInput float64

	fmt.Print(text)
	fmt.Scan(&userInput)

	if userInput < 0 {
		fmt.Println("Error: Input must be a non-negative number.")
		os.Exit(1)
	}

	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func saveFinancialsToFile(ebt, profit, ratio float64) {
	financialsString := fmt.Sprintf("EBT: %.1f\nNet Income: %.1f\nEBT to Net Income Ratio: %.2f\n", ebt, profit, ratio)
	os.WriteFile("Financials", []byte(financialsString), 0644)
}
