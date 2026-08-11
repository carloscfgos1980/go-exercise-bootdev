package functions

import (
	"fmt"
	"testing"
)

func TestGetMonthlyPrice(t *testing.T) {
	tests := []struct {
		tier     string
		expected int
	}{
		{"basic", 10000},
		{"premium", 15000},
		{"enterprise", 50000},
		{"unknown", 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("tier=%s", test.tier), func(t *testing.T) {
			result := GetMonthlyPrice(test.tier)
			if result != test.expected {
				t.Errorf("GetMonthlyPrice(%q) = %d; want %d", test.tier, result, test.expected)
			}
		})
	}
}

func TestMonthlyBillIncrease(t *testing.T) {
	tests := []struct {
		costPerSend  int
		numLastMonth int
		numThisMonth int
		expected     int
	}{
		{100, 10, 15, 500},
		{200, 5, 10, 1000},
		{50, 20, 25, 250},
		{0, 10, 20, 0},
		{100, 10, 0, -1000},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("costPerSend=%d,numLastMonth=%d,numThisMonth=%d", test.costPerSend, test.numLastMonth, test.numThisMonth), func(t *testing.T) {
			result := MonthlyBillIncrease(test.costPerSend, test.numLastMonth, test.numThisMonth)
			if result != test.expected {
				t.Errorf("MonthlyBillIncrease(%d, %d, %d) = %d; want %d", test.costPerSend, test.numLastMonth, test.numThisMonth, result, test.expected)
			}
		})
	}
}

func TestGetProductMessage(t *testing.T) {
	tests := []struct {
		tier     string
		expected string
	}{
		{"basic", "You get 1,000 texts per month for $30 per month."},
		{"premium", "You get 50,000 texts per month for $60 per month."},
		{"enterprise", "You get unlimited texts per month for $100 per month."},
		{"unknown", "You get  for ."},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("tier=%s", test.tier), func(t *testing.T) {
			result := GetProductMessage(test.tier)
			if result != test.expected {
				t.Errorf("GetProductMessage(%q) = %q; want %q", test.tier, result, test.expected)
			}
		})
	}
}

func TestYearsUntilEvents(t *testing.T) {
	tests := []struct {
		age                         int
		expectedYearsUntilAdult     int
		expectedYearsUntilDrinking  int
		expectedYearsUntilCarRental int
	}{
		{17, 1, 4, 8},
		{18, 0, 3, 7},
		{20, 0, 1, 5},
		{21, 0, 0, 4},
		{24, 0, 0, 1},
		{25, 0, 0, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("age=%d", test.age), func(t *testing.T) {
			resultAdult, resultDrinking, resultCarRental := YearsUntilEvents(test.age)
			if resultAdult != test.expectedYearsUntilAdult || resultDrinking != test.expectedYearsUntilDrinking || resultCarRental != test.expectedYearsUntilCarRental {
				t.Errorf("YearsUntilEvents(%d) = (%d, %d, %d); want (%d, %d, %d)", test.age, resultAdult, resultDrinking, resultCarRental, test.expectedYearsUntilAdult, test.expectedYearsUntilDrinking, test.expectedYearsUntilCarRental)
			}
		})
	}
}

func TestReformat(t *testing.T) {
	tests := []struct {
		message       string
		formatter     func(string) string
		formatterName string
		expected      string
	}{
		{"hello", addExclamation, "addExclamation", "TEXTIO: hello!!!"},
		{"hello there", addPeriod, "addPeriod", "TEXTIO: hello there..."},
		{"moor der ehT", reverseString, "reverseString", "TEXTIO: The red room"},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("message=%s,formatter=%s", test.message, test.formatterName), func(t *testing.T) {
			result := Reformat(test.message, test.formatter)
			if result != test.expected {
				t.Errorf("Reformat(%q, %s) = %q; want %q", test.message, test.formatterName, result, test.expected)
			}
		})
	}
}

func addPeriod(s string) string {
	return s + "."
}

func addExclamation(s string) string {
	return s + "!"
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func TestSplitEmail(t *testing.T) {
	tests := []struct {
		email    string
		username string
		domain   string
	}{
		{"drogon@dragonstone.com", "drogon", "dragonstone.com"},
		{"rhaenyra@targaryen.com", "rhaenyra", "targaryen.com"},
		{"viserys@kingslanding.com", "viserys", "kingslanding.com"},
		{"aegon@stormsend.com", "aegon", "stormsend.com"},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("email=%s", test.email), func(t *testing.T) {
			username, domain := SplitEmail(test.email)
			if username != test.username || domain != test.domain {
				t.Errorf("SplitEmail(%q) = (%q, %q); want (%q, %q)", test.email, username, domain, test.username, test.domain)
			}
		})
	}
}
