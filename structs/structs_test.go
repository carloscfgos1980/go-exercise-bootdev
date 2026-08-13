package structs

import "testing"

func TestCanSendMessage(t *testing.T) {
	tests := []struct {
		mToSend  MessageToSend
		expected bool
	}{
		{MessageToSend{
			Message:   "you have an appointment tomorrow",
			Sender:    User{Name: "Brenda Halafax", Number: 16545550987},
			Recipient: User{Name: "Sally Sue", Number: 19035558973},
		}, true},
		{MessageToSend{
			Message:   "you have an event tomorrow",
			Sender:    User{Number: 16545550987},
			Recipient: User{Name: "Suzie Sall", Number: 19035558973},
		}, false},
		{MessageToSend{
			Message:   "you have an birthday tomorrow",
			Sender:    User{Name: "Jason Bjorn", Number: 16545550987},
			Recipient: User{Name: "Jim Bond"},
		}, false},
		{MessageToSend{
			Message:   "you have a party tomorrow",
			Sender:    User{Name: "Njorn Halafax"},
			Recipient: User{Name: "Becky Sue", Number: 19035558973},
		}, false},
		{MessageToSend{
			Message:   "you have a birthday tomorrow",
			Sender:    User{Name: "Eli Halafax", Number: 16545550987},
			Recipient: User{Number: 19035558973},
		}, false},
	}
	for _, test := range tests {
		result := CanSendMessage(test.mToSend)
		if result != test.expected {
			t.Errorf("For message %+v, expected %v, got %v", test.mToSend, test.expected, result)
		}
	}
}
