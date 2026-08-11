package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/carloscfgos1980/go-exercise-bootdev/condictionals"
)

func main() {
	commands := map[string]func(){
		"calculate-balance": condictionals.CalculateBalance,
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
