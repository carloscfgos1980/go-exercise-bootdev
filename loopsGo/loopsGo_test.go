package loopsGo

import (
	"fmt"
	"testing"
)

func TestBulkSend(t *testing.T) {
	tests := []struct {
		numMessages int
		expected    float64
	}{
		{10, 10.45},
		{20, 21.9},
		{0, 0.0},
		{1, 1.0},
		{5, 5.10},
		{30, 34.35},
	}
	for _, tt := range tests {
		result := BulkSend(tt.numMessages)
		roundResult := float64(int(result*100)) / 100        // Round to 2 decimal places
		roundExpected := float64(int(tt.expected*100)) / 100 // Round to 2 decimal places
		fmt.Printf("BulkSend(%d) = %v; want %v\n", tt.numMessages, roundResult, roundExpected)
		if roundResult != roundExpected {
			t.Errorf("BulkSend(%d) = %v; want %v", tt.numMessages, roundResult, roundExpected)
		}
	}
}

func TestGetMaxMessagesToSend(t *testing.T) {
	tests := []struct {
		costMultiplier  float64
		budgetInPennies int
		expected        int
	}{
		{1.1, 5, 4},
		{1.3, 10, 5},
		{1.35, 25, 7},
		{1.2, 1, 1},
		{1.2, 15, 7},
		{1.3, 20, 7},
		{1.5, 0, 0},
	}
	for _, tt := range tests {
		result := GetMaxMessagesToSend(tt.costMultiplier, tt.budgetInPennies)
		fmt.Printf("GetMaxMessagesToSend(%v, %v) = %v; want %v\n", tt.costMultiplier, tt.budgetInPennies, result, tt.expected)
		if result != tt.expected {
			t.Errorf("GetMaxMessagesToSend(%v, %v) = %v; want %v", tt.costMultiplier, tt.budgetInPennies, result, tt.expected)
		}
	}
}

func TestCountConnections(t *testing.T) {
	tests := []struct {
		groupSize int
		expected  int
	}{
		{1, 0},
		{2, 1},
		{3, 3},
		{4, 6},
		{0, 0},
		{10, 45},
		{100, 4950},
		{1000, 499500},
	}
	for _, tt := range tests {
		result := CountConnections(tt.groupSize)
		fmt.Printf("CountConnections(%v) = %v; want %v\n", tt.groupSize, result, tt.expected)
		if result != tt.expected {
			t.Errorf("CountConnections(%v) = %v; want %v", tt.groupSize, result, tt.expected)
		}
	}
}
