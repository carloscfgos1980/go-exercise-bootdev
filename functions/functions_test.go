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
