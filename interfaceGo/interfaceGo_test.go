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

func TestSendMessage1(t *testing.T) {
	tests := []struct {
		exp          expense1
		format       formatter1
		expectedText string
		expectedCost int
	}{
		{
			exp:          email1{isSubscribed: true, body: "Hello"},
			format:       email1{isSubscribed: true, body: "Hello"},
			expectedText: "'Hello' | Subscribed",
			expectedCost: 10,
		},
		{
			exp:          email1{isSubscribed: false, body: "Hello"},
			format:       email1{isSubscribed: false, body: "Hello"},
			expectedText: "'Hello' | Not Subscribed",
			expectedCost: 25,
		},
	}
	for _, test := range tests {
		resultText, resultCost := SendMessage1(test.exp, test.format)
		fmt.Printf("Testing expense %+v, got text %q and cost %d\n", test.exp, resultText, resultCost)
		if resultText != test.expectedText || resultCost != test.expectedCost {
			t.Errorf("For expense %+v, expected text %q and cost %d, got text %q and cost %d", test.exp, test.expectedText, test.expectedCost, resultText, resultCost)
		}
	}
}

func TestGetExpenseReport(t *testing.T) {
	tests := []struct {
		exp          expense
		expectedText string
		expectedCost float64
	}{
		{
			email{
				isSubscribed: true,
				body:         "Whoa there!",
				toAddress:    "soldier@monty.com"},
			"soldier@monty.com",
			0.11,
		},
		{
			sms{
				isSubscribed:  false,
				body:          "Halt! Who goes there?",
				toPhoneNumber: "+155555509832",
			},
			"+155555509832",
			2.1,
		},
		{
			email{
				isSubscribed: false,
				body:         "It is I, Arthur, son of Uther Pendragon, from the castle of Camelot. King of the Britons, defeator of the Saxons, sovereign of all England!",
				toAddress:    "soldier@monty.com",
			},
			"soldier@monty.com",
			6.95,
		},
		{
			email{
				isSubscribed: true,
				body:         "Pull the other one!",
				toAddress:    "arthur@monty.com",
			},
			"arthur@monty.com",
			0.19,
		},
		{
			sms{
				isSubscribed:  true,
				body:          "I am. And this my trusty servant Patsy.",
				toPhoneNumber: "+155555509832",
			},
			"+155555509832",
			1.17,
		},
		{
			invalid{},
			"",
			0.0,
		},
	}
	for _, test := range tests {
		resultText, resultCost := GetExpenseReport(test.exp)
		fmt.Printf("Testing expense %+v, got text %q and cost %f\n", test.exp, resultText, resultCost)
		if resultText != test.expectedText || resultCost != test.expectedCost {
			t.Errorf("For expense %+v, expected text %q and cost %f, got text %q and cost %f", test.exp, test.expectedText, test.expectedCost, resultText, resultCost)
		}
	}
}

func TestGetExpenseReport2(t *testing.T) {
	tests := []struct {
		exp          expense
		expectedText string
		expectedCost float64
	}{
		{
			email{
				isSubscribed: true,
				body:         "Whoa there!",
				toAddress:    "soldier@monty.com"},
			"soldier@monty.com",
			0.11,
		},
		{
			sms{
				isSubscribed:  false,
				body:          "Halt! Who goes there?",
				toPhoneNumber: "+155555509832",
			},
			"+155555509832",
			2.1,
		},
		{
			email{
				isSubscribed: false,
				body:         "It is I, Arthur, son of Uther Pendragon, from the castle of Camelot. King of the Britons, defeator of the Saxons, sovereign of all England!",
				toAddress:    "soldier@monty.com",
			},
			"soldier@monty.com",
			6.95,
		},
		{
			email{
				isSubscribed: true,
				body:         "Pull the other one!",
				toAddress:    "arthur@monty.com",
			},
			"arthur@monty.com",
			0.19,
		},
		{
			sms{
				isSubscribed:  true,
				body:          "I am. And this my trusty servant Patsy.",
				toPhoneNumber: "+155555509832",
			},
			"+155555509832",
			1.17,
		},
		{
			invalid{},
			"",
			0.0,
		},
	}
	for _, test := range tests {
		resultText, resultCost := GetExpenseReport(test.exp)
		fmt.Printf("Testing expense %+v, got text %q and cost %f\n", test.exp, resultText, resultCost)
		if resultText != test.expectedText || resultCost != test.expectedCost {
			t.Errorf("For expense %+v, expected text %q and cost %f, got text %q and cost %f", test.exp, test.expectedText, test.expectedCost, resultText, resultCost)
		}
	}
}

func TestSendMessage2(t *testing.T) {
	tests := []struct {
		format   formatter
		expected string
	}{
		{plainText{message: "Hello, World!"}, "Hello, World!"},
		{bold{message: "Bold Message"}, "**Bold Message**"},
		{code{message: "Code Message"}, "`Code Message`"},
	}
	for _, test := range tests {
		result := SendMessage2(test.format)
		fmt.Printf("Testing format %+v, got %q\n", test.format, result)
		if result != test.expected {
			t.Errorf("For format %+v, expected %q, got %q", test.format, test.expected, result)
		}
	}
}
