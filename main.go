package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/carloscfgos1980/go-exercise-bootdev/condictionals"
	"github.com/carloscfgos1980/go-exercise-bootdev/functions"
)

func main() {
	commands := map[string]func(){
		"calculate-balance": condictionals.CalculateBalance,
		"monthly-price": func() {
			if len(os.Args) < 3 {
				fmt.Printf("usage: go run . monthly-price <tier>\n")
				return
			}
			fmt.Println(functions.GetMonthlyPrice(os.Args[2]))
		},
		"monthly-bill-increase": func() {
			costPerSendInt := 1000
			numLastMonthInt := 10
			numThisMonthInt := 15

			result := functions.MonthlyBillIncrease(costPerSendInt, numLastMonthInt, numThisMonthInt)
			fmt.Println(result)
		},
		"product-message": func() {
			if len(os.Args) < 3 {
				fmt.Printf("usage: go run . product-message <tier>\n")
				return
			}
			fmt.Println(functions.GetProductMessage(os.Args[2]))
		},
		"years-until-events": func() {
			if len(os.Args) < 3 {
				fmt.Printf("usage: go run . years-until-events <age>\n")
				return
			}
			age := 0
			fmt.Sscanf(os.Args[2], "%d", &age)
			yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := functions.YearsUntilEvents(age)
			fmt.Printf("Years until adult: %d\n", yearsUntilAdult)
			fmt.Printf("Years until drinking: %d\n", yearsUntilDrinking)
			fmt.Printf("Years until car rental: %d\n", yearsUntilCarRental)
		},
		"print-reports": func() {
			intro := "Welcome to the monthly report"
			body := "This month we had a 10% increase in sales"
			outro := "Thank you for your business"
			functions.PrintReports(intro, body, outro)
		},
		"bootup": functions.Bootup,
		"split-email": func() {
			if len(os.Args) < 3 {
				fmt.Printf("usage: go run . split-email <email>\n")
				return
			}
			email := os.Args[2]
			username, domain := functions.SplitEmail(email)
			fmt.Printf("Username: %s\n", username)
			fmt.Printf("Domain: %s\n", domain)
		},
		"place-order": func() {
			if len(os.Args) < 5 {
				fmt.Printf("usage: go run . place-order <productID> <quantity> <accountBalance>\n")
				return
			}
			productID := os.Args[2]
			quantity := 0
			accountBalance := 0.0
			fmt.Sscanf(os.Args[3], "%d", &quantity)
			fmt.Sscanf(os.Args[4], "%f", &accountBalance)
			success, remainingBalance := functions.PlaceOrder(productID, quantity, accountBalance)
			fmt.Printf("Success: %v\n", success)
			fmt.Printf("Remaining balance: %.2f\n", remainingBalance)
		},
	}

	if len(os.Args) < 2 {
		fmt.Printf("usage: go run . [%s]\n", strings.Join(commandNames(commands), "|"))
		return
	}

	run, ok := commands[os.Args[1]]
	if !ok {
		fmt.Printf("unknown command %q. available: %s\n", os.Args[1], strings.Join(commandNames(commands), ", "))
		return
	}

	run()
}

func commandNames(commands map[string]func()) []string {
	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
