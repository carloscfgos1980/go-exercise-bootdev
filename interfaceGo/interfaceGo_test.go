package interfaceGo

import (
	"fmt"
	"testing"
	"time"
)

func TestSendMessage(t *testing.T) {
	tests := []struct {
		msg          message
		expectedText string
		expectedCost int
	}{
		{
			msg:          birthdayMessage{time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC), "John Doe"},
			expectedText: "Hi John Doe, it is your birthday on 1994-03-21T00:00:00Z",
			expectedCost: 168,
		},
		{
			msg:          sendingReport{"First Report", 10},
			expectedText: `Your "First Report" report is ready. You've sent 10 messages.`,
			expectedCost: 183,
		},
		{
			msg:          birthdayMessage{time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC), "Bill Deer"},
			expectedText: "Hi Bill Deer, it is your birthday on 1934-05-01T00:00:00Z",
			expectedCost: 171,
		},
		{
			msg:          sendingReport{"Second Report", 20},
			expectedText: `Your "Second Report" report is ready. You've sent 20 messages.`,
			expectedCost: 186,
		},
	}
	for _, test := range tests {
		resultText, resultCost := SendMessage(test.msg)
		fmt.Printf("Testing message %+v, got text %q and cost %d\n", test.msg, resultText, resultCost)
		if resultText != test.expectedText || resultCost != test.expectedCost {
			t.Errorf("For message %+v, expected text %q and cost %d, got text %q and cost %d", test.msg, test.expectedText, test.expectedCost, resultText, resultCost)
		}
	}
}
