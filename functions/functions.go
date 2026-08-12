package functions

import "fmt"

// using unit testing to validate the functions
func GetMonthlyPrice(tier string) int {
	switch tier {
	case "basic":
		return 10000
	case "premium":
		return 15000
	case "enterprise":
		return 50000
	default:
		return 0
	}
}

// Passing variable per value
func MonthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	var lastMonthBill int
	var thisMonthBill int
	lastMonthBill = getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill = getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	return costPerSend * messagesSent
}

// Ignoring return values
func GetProductMessage(tier string) string {
	quantityMsg, priceMsg, _ := getProductInfo(tier)
	return "You get " + quantityMsg + " for " + priceMsg + "."
}

func getProductInfo(tier string) (string, string, string) {
	switch tier {
	case "basic":
		return "1,000 texts per month", "$30 per month", "most popular"
	case "premium":
		return "50,000 texts per month", "$60 per month", "best value"
	case "enterprise":
		return "unlimited texts per month", "$100 per month", "customizable"
	default:
		return "", "", ""
	}
}

// Named return values
func YearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	// don't touch below this line

	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}

// Explicit return values
func YearsUntilEvents2(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}

// functions as values
func Reformat(message string, formatter func(string) string) string {
	once := formatter(message)
	twice := formatter(once)
	thrice := formatter(twice)
	prefix := "TEXTIO: "
	return prefix + thrice
}

// Anonymous functions
func PrintReports(intro, body, outro string) {
	printCostReport(func(intro string) int {
		return len(intro) * 2
	}, intro)
	printCostReport(func(body string) int {
		return len(intro) * 3
	}, body)
	printCostReport(func(outro string) int {
		return len(intro) * 4
	}, outro)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}

// Defer statement
func Bootup() {
	defer fmt.Println("TEXTIO BOOTUP DONE")
	ok := connectToDB()
	if !ok {
		return
	}
	ok = connectToPaymentProvider()
	if !ok {
		return
	}
	fmt.Println("All systems ready!")
}

var shouldConnectToDB = true

func connectToDB() bool {
	fmt.Println("Connecting to database...")
	if shouldConnectToDB {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

var shouldConnectToPaymentProvider = true

func connectToPaymentProvider() bool {
	fmt.Println("Connecting to payment provider...")
	if shouldConnectToPaymentProvider {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

// Block scope
func SplitEmail(email string) (string, string) {
	username, domain := "", ""

	for i, r := range email {
		if r == '@' {
			username = email[:i]
			domain = email[i+1:]
			break
		}
	}
	return username, domain
}

// Processing orders
func PlaceOrder(productID string, quantity int, accountBalance float64) (bool, float64) {
	availableStock := amountInStock(productID)
	if quantity > availableStock {
		return false, accountBalance
	}

	totalCost := calcPrice(productID, quantity)
	if totalCost > accountBalance {
		return false, accountBalance
	}

	remainingBalance := accountBalance - totalCost
	return true, remainingBalance
}

func calcPrice(productID string, quantity int) float64 {
	return priceList(productID) * float64(quantity)
}

func priceList(productID string) float64 {
	switch productID {
	case "1":
		return 1.50
	case "2":
		return 2.25
	case "3":
		return 3.00
	case "4":
		return 1.00
	case "5":
		return 2.50
	case "6":
		return 8.99
	case "7":
		return 22.50
	case "8":
		return 50.00
	case "9":
		return 999.99
	default:
		return 0.00
	}
}

func amountInStock(productID string) int {
	switch productID {
	case "1":
		return 11
	case "2":
		return 25
	case "3":
		return 4
	case "4":
		return 6
	case "5":
		return 50
	case "6":
		return 2
	case "7":
		return 0
	case "8":
		return 99
	case "9":
		return 1
	default:
		return 0
	}
}

// Closures
func Adder() func(int) int {
	sum := 0
	return func(value int) int {
		sum += value
		return sum
	}
}
