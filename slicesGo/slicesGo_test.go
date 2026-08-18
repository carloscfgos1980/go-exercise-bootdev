package slicesgo

import (
	"fmt"
	"testing"
)

func TestGetMessageWithRetries(t *testing.T) {
	tests := []struct {
		messages         []string
		expectedMessages [3]string
		expectedCosts    [3]int
	}{
		{
			[]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[3]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[3]int{46, 92, 119},
		},
		{
			[]string{"It's the spring fling sale!", "Don't miss this event!", "Last chance."},
			[3]string{"It's the spring fling sale!", "Don't miss this event!", "Last chance."},
			[3]int{27, 49, 61},
		},
	}
	for _, tt := range tests {
		resultMessages, resultCosts := GetMessageWithRetries(tt.messages[0], tt.messages[1], tt.messages[2])
		fmt.Printf("GetMessageWithRetries(%v) = %v, %v; want %v, %v\n", tt.messages, resultMessages, resultCosts, tt.expectedMessages, tt.expectedCosts)
		if resultMessages != tt.expectedMessages || resultCosts != tt.expectedCosts {
			t.Errorf("GetMessageWithRetries(%v) = %v, %v; want %v, %v", tt.messages, resultMessages, resultCosts, tt.expectedMessages, tt.expectedCosts)
		}
	}
}

func TestGetMessageWithRetriesForPlan(t *testing.T) {
	tests := []struct {
		plan             string
		messages         [3]string
		expectedMessages []string
		expectedErr      error
	}{
		{
			planFree,
			[3]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[]string{"Hello sir/madam can I interest you in a yacht?", "Please I'll even give you an Amazon gift card?"},
			nil,
		},
		{
			planPro,
			[3]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			[]string{
				"Hello sir/madam can I interest you in a yacht?",
				"Please I'll even give you an Amazon gift card?",
				"You're missing out big time",
			},
			nil,
		},
		{
			"invalid plan",
			[3]string{
				"You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?",
				"Wouldn't you?",
				"Wouldn't you???",
			},
			nil,
			fmt.Errorf("unsupported plan"),
		},
	}
	for _, tt := range tests {
		resultMessages, resultErr := GetMessageWithRetriesForPlan(tt.plan, tt.messages)
		fmt.Printf("GetMessageWithRetriesForPlan(%v, %v) = %v, %v; want %v, %v\n", tt.plan, tt.messages, resultMessages, resultErr, tt.expectedMessages, tt.expectedErr)
		if !equalSlices(resultMessages, tt.expectedMessages) || !equalErrors(resultErr, tt.expectedErr) {
			t.Errorf("GetMessageWithRetriesForPlan(%v, %v) = %v, %v; want %v, %v", tt.plan, tt.messages, resultMessages, resultErr, tt.expectedMessages, tt.expectedErr)
		}
	}

}

func equalSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalErrors(a, b error) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Error() == b.Error()
}

func TestGetMessageCosts(t *testing.T) {
	tests := []struct {
		messages    []string
		expected    []float64
		expectedCap int
	}{
		{
			[]string{"Welcome to the movies!", "Enjoy your popcorn!"},
			[]float64{0.22, 0.19},
			2,
		},
		{
			[]string{"I don't want to be here anymore", "Can we go home?", "I'm hungry", "I'm bored"},
			[]float64{0.31, 0.15, 0.1, 0.09},
			4,
		},
		{[]string{}, []float64{}, 0},
		{[]string{""}, []float64{0}, 1},
		{[]string{"Hello", "Hi", "Hey"}, []float64{0.05, 0.02, 0.03}, 3},
	}
	for _, tt := range tests {
		cost := GetMessageCosts(tt.messages)
		fmt.Printf("GetMessageCosts(%v) = %v; want %v; want capacity = %v, got %v\n", tt.messages, cost, tt.expected, tt.expectedCap, cap(cost))
		if !equalSlices(tt.expected, cost) || cap(cost) != tt.expectedCap {
			t.Errorf("GetMessageCosts(%v) = %v; want %v; cap = %v; want cap = %v", tt.messages, cost, tt.expected, cap(cost), tt.expectedCap)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 120},
		{[]int{}, 0},
		{[]int{5}, 5},
	}
	for _, tt := range tests {
		result := Sum(tt.nums...)
		fmt.Printf("Sum(%v) = %v; want %v\n", tt.nums, result, tt.expected)
		if result != tt.expected {
			t.Errorf("Sum(%v) = %v; want %v", tt.nums, result, tt.expected)
		}
	}
}
