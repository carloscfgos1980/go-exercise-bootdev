package generics

import (
	"fmt"
	"testing"
	"time"
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

func TestChargeForLineItem(t *testing.T) {
	tests := []struct {
		newItem           lineItem
		oldItems          []lineItem
		balance           float64
		expected          []lineItem
		expectedBalance   float64
		expectedErrString string
	}{
		{
			newItem: subscription{
				userEmail: "geralt@rivia.com",
				startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				interval:  "yearly",
			},
			oldItems: []lineItem{
				subscription{
					userEmail: "yen@vengerberg.com",
					startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					interval:  "monthly",
				},
				oneTimeUsagePlan{
					userEmail:        "triss@maribor",
					numEmailsAllowed: 100,
				},
			},
			balance: 1000.00,
			expected: []lineItem{
				subscription{
					userEmail: "yen@vengerberg.com",
					startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					interval:  "monthly",
				},
				oneTimeUsagePlan{
					userEmail:        "triss@maribor",
					numEmailsAllowed: 100,
				},
				subscription{
					userEmail: "geralt@rivia.com",
					startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					interval:  "yearly",
				},
			},
			expectedBalance:   750.00,
			expectedErrString: "",
		},
		{
			newItem: subscription{
				userEmail: "geralt@rivia.com",
				startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				interval:  "yearly",
			},
			oldItems: []lineItem{
				subscription{
					userEmail: "yen@vengerberg.com",
					startDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					interval:  "monthly",
				},
				oneTimeUsagePlan{
					userEmail:        "triss@maribor",
					numEmailsAllowed: 100,
				},
			},
			balance:           200.00,
			expected:          nil,
			expectedBalance:   0.0,
			expectedErrString: "insufficient funds",
		},
	}
	for _, tt := range tests {
		got, gotBalance, err := ChargeForLineItem(tt.newItem, tt.oldItems, tt.balance)
		if err != nil && err.Error() != tt.expectedErrString {
			t.Errorf("ChargeForLineItem() error = %v; want %v", err.Error(), tt.expectedErrString)
		}
		if !equalLineItems(got, tt.expected) {
			t.Errorf("ChargeForLineItem() got = %v; want %v", got, tt.expected)
		}
		if gotBalance != tt.expectedBalance {
			t.Errorf("ChargeForLineItem() gotBalance = %v; want %v", gotBalance, tt.expectedBalance)
		}
	}
}

func equalLineItems(a, b []lineItem) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].GetName() != b[i].GetName() || a[i].GetCost() != b[i].GetCost() {
			return false
		}
	}
	return true
}
