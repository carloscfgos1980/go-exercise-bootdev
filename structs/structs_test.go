package structs

import (
	"fmt"
	"testing"
)

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

func TestGetBasicAuth(t *testing.T) {
	tests := []struct {
		auth     authenticationInfo
		expected string
	}{
		{authenticationInfo{"Google", "12345"}, "Authorization: Basic Google:12345"},
		{authenticationInfo{"Bing", "98765"}, "Authorization: Basic Bing:98765"},
		{authenticationInfo{"DDG", "76921"}, "Authorization: Basic DDG:76921"},
	}
	for _, test := range tests {
		result := test.auth.GetBasicAuth()
		fmt.Printf("Testing auth %+v, got %v\n", test.auth, result)
		if result != test.expected {
			t.Errorf("For auth %+v, expected %v, got %v", test.auth, test.expected, result)
		}
	}
}

func TestNewUser(t *testing.T) {
	tests := []struct {
		name           string
		membershipType string
	}{
		{"Syl", "standard"},
		{"Pattern", "premium"},
		{"Pattern", "standard"},
		{"Renarin", "standard"},
		{"Lift", "premium"},
		{"Dalinar", "standard"},
	}
	for _, test := range tests {
		user := NewUser(test.name, test.membershipType)
		fmt.Printf("Testing user %+v, got membership %+v\n", user, user.Membership)
		if user.Name != test.name {
			t.Errorf("For name %s, expected %s, got %s", test.name, test.name, user.Name)
		}
		if user.Membership.Type != test.membershipType {
			t.Errorf("For membership type %s, expected %s, got %s", test.membershipType, test.membershipType, user.Membership.Type)
		}
		if test.membershipType == "premium" && user.Membership.MessageCharLimit != 1000 {
			t.Errorf("For premium membership, expected char limit 1000, got %d", user.Membership.MessageCharLimit)
		}
		if test.membershipType != "premium" && user.Membership.MessageCharLimit != 100 {
			t.Errorf("For non-premium membership, expected char limit 100, got %d", user.Membership.MessageCharLimit)
		}
	}
}
