package functions

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
