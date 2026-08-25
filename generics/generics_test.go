package generics

import (
	"fmt"
	"testing"
)

func TestGetLast(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected interface{}
	}{
		{[]int{}, 0},
		{[]bool{true, false, true, true, false}, false},
		{[]int{1, 2, 3, 4}, 4},
		{[]string{"a", "b", "c", "d"}, "d"},
	}
	for _, tt := range tests {
		switch input := tt.input.(type) {
		case []int:
			got := getLast(input)
			if got != tt.expected {
				fmt.Printf("getLast(%v) = %v; want %v\n", input, got, tt.expected)
				t.Errorf("getLast(%v) = %v; want %v", input, got, tt.expected)
			}
		case []bool:
			got := getLast(input)
			fmt.Printf("getLast(%v) = %v; want %v\n", input, got, tt.expected)
			if got != tt.expected {
				t.Errorf("getLast(%v) = %v; want %v", input, got, tt.expected)
			}
		case []string:
			got := getLast(input)
			fmt.Printf("getLast(%v) = %v; want %v\n", input, got, tt.expected)
			if got != tt.expected {
				t.Errorf("getLast(%v) = %v; want %v", input, got, tt.expected)
			}
		default:
			t.Errorf("unsupported type: %T", input)
		}
	}
}
