package loopsGo

import "fmt"

// loops in go
func BulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (float64(i) * 0.01)
	}
	return totalCost
}

// There Is No While Loop in Go

func GetMaxMessagesToSend(costMultiplier float64, budgetInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(budgetInPennies) - actualCostInPennies
	for balance > 0 {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}

// Fizzbuzz
func Fizzbuzz() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
