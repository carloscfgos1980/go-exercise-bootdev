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
