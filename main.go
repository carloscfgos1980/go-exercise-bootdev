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
