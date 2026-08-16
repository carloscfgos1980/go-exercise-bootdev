package loops

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
