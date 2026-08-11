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
