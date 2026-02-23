package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	outputText("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected annual return rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the investment duration in years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// formattedFV := fmt.Sprintf("Future Value: %.0f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Future Value (Adjusted for inflation): %.0f\n", futureRealValue)
	// fmt.Println("Future Value:", futureValue)
	fmt.Printf("Future Value: %.0f\nFuture Value (Adjusted for inflation): %.0f\n", futureValue, futureRealValue)
	// fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureRealValue
}
