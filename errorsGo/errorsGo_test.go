package errorsgo

import (
	"errors"
	"fmt"
	"testing"
)

func TestSendSMSToCouple(t *testing.T) {
	tests := []struct {
		msgToCustomer string
		msgToSpouse   string
		expectedCost  int
		expectedErr   error
	}{
		{"Thanks for coming in to our flower shop today!", "We hope you enjoyed your gift.", 0, fmt.Errorf("can't send texts over 25 characters")},
		{"Thanks for joining us!", "Have a good day.", 76, nil},
		{"Thank you.", "Enjoy!", 32, nil},
		{"We loved having you in!", "We hope the rest of your evening is fantastic.", 0, fmt.Errorf("can't send texts over 25 characters")},
	}
	for _, tt := range tests {
		cost, err := SendSMSToCouple(tt.msgToCustomer, tt.msgToSpouse)
		fmt.Printf("Testing SendSMSToCouple(%q, %q), got cost %v and error %v\n", tt.msgToCustomer, tt.msgToSpouse, cost, err)
		if cost != tt.expectedCost || (err != nil && err.Error() != tt.expectedErr.Error()) {
			t.Errorf("SendSMSToCouple(%q, %q) = %v, %v; want %v, %v",
				tt.msgToCustomer, tt.msgToSpouse, cost, err, tt.expectedCost, tt.expectedErr)
		}
	}
}

func TestGetSMSErrorString(t *testing.T) {
	tests := []struct {
		cost      float64
		recipient string
		expected  string
	}{
		{1.4, "+1 (435) 555 0923", "SMS that costs $1.40 to be sent to '+1 (435) 555 0923' cannot be sent"},
		{2.1, "+2 (702) 555 3452", "SMS that costs $2.10 to be sent to '+2 (702) 555 3452' cannot be sent"},
		{32.1, "+1 (801) 555 7456", "SMS that costs $32.10 to be sent to '+1 (801) 555 7456' cannot be sent"},
		{14.4, "+1 (234) 555 6545", "SMS that costs $14.40 to be sent to '+1 (234) 555 6545' cannot be sent"},
	}
	for _, tt := range tests {
		result := GetSMSErrorString(tt.cost, tt.recipient)
		fmt.Printf("Testing GetSMSErrorString(%.2f, %q), got result %q\n", tt.cost, tt.recipient, result)
		if result != tt.expected {
			t.Errorf("GetSMSErrorString(%.2f, %q) = %q; want %q", tt.cost, tt.recipient, result, tt.expected)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		dividend, divisor, expected float64
		expectedError               string
	}{
		{10, 2, 5, ""},
		{15, 3, 5, ""},
		{10, 0, 0, "cannot divide 10 by zero"},
		{15, 0, 0, "cannot divide 15 by zero"},
	}
	for _, tt := range tests {
		result, err := Divide(tt.dividend, tt.divisor)
		fmt.Printf("Testing Divide(%v, %v), got result %v and error %v\n", tt.dividend, tt.divisor, result, err)
		if result != tt.expected || (err != nil && err.Error() != tt.expectedError) {
			t.Errorf("Divide(%v, %v) = %v, %v; want %v, %v", tt.dividend, tt.divisor, result, err, tt.expected, tt.expectedError)
		}
	}
}

func TestDivide2(t *testing.T) {
	tests := []struct {
		x, y, expected float64
		expectedErr    string
	}{
		{10, 0, 0, "no dividing by 0"},
		{10, 2, 5, ""},
		{15, 30, 0.5, ""},
		{6, 3, 2, ""},
	}
	for _, tt := range tests {
		result, err := Divide2(tt.x, tt.y)
		fmt.Printf("Testing Divide2(%v, %v), got result %v and error %v\n", tt.x, tt.y, result, err)
		if result != tt.expected || (err != nil && err.Error() != tt.expectedErr) {
			t.Errorf("Divide2(%v, %v) = %v, %v; want %v, %v", tt.x, tt.y, result, err, tt.expected, tt.expectedErr)
		}
	}
}

func TestValidStatus(t *testing.T) {
	tests := []struct {
		status      string
		expectedErr error
	}{
		{"", errors.New("status cannot be empty")},
		{"This is a valid status update that is well within the character limit.", nil},
		{"This status update is way too long. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco.", errors.New("status exceeds 140 characters")},
	}
	for _, tt := range tests {
		err := ValidateStatus(tt.status)
		fmt.Printf("Testing ValidateStatus(%q), got error %v\n", tt.status, err)
		if (err != nil && tt.expectedErr == nil) || (err == nil && tt.expectedErr != nil) || (err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error()) {
			t.Errorf("ValidateStatus(%q) = %v; want %v", tt.status, err, tt.expectedErr)
		}
	}
}
