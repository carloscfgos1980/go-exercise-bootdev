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
