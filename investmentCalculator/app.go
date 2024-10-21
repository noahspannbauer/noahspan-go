package main

import "fmt"
import "math"

const inflationRate = 2.5
var investmentAmount float64
var years float64
var expectedReturnRate float64

func main() {
	

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)


	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (adjusted for inflaction):", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
} 

func calculateFutureValues (investmentAmount, expectedFutureValue, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	rfv = fv / math.Pow(1 + inflationRate/100, years)

	return fv, rfv
}