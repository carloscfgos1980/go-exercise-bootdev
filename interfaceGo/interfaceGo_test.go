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

func TestGetEmployeeSalary(t *testing.T) {
	tests := []struct {
		emp            employee
		expectedSalary int
	}{
		{emp: fullTime{name: "Bob", salary: 7300}, expectedSalary: 7300},
		{emp: contractor{name: "Jill", hourlyPay: 872, hoursPerYear: 982}, expectedSalary: 856304},
		{emp: fullTime{name: "Alice", salary: 10000}, expectedSalary: 10000},
		{emp: contractor{name: "John", hourlyPay: 1000, hoursPerYear: 1000}, expectedSalary: 1000000},
	}
	for _, test := range tests {
		resultSalary := GetEmployeeSalary(test.emp)
		fmt.Printf("Testing employee %+v, got salary %d\n", test.emp, resultSalary)
		if resultSalary != test.expectedSalary {
			t.Errorf("For employee %+v, expected salary %d, got %d", test.emp, test.expectedSalary, resultSalary)
		}
	}
}
